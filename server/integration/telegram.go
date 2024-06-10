package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"time"
	"unicode/utf16"

	"github.com/lithammer/shortuuid/v4"
	"github.com/pkg/errors"
	"github.com/yourselfhosted/gomark/ast"
	"github.com/yourselfhosted/gomark/parser"
	"github.com/yourselfhosted/gomark/parser/tokenizer"

	"github.com/Syuq/Locket/plugin/telegram"
	"github.com/Syuq/Locket/plugin/webhook"
	storepb "github.com/Syuq/Locket/proto/gen/store"
	apiv1 "github.com/Syuq/Locket/server/route/api/v1"
	apiv2 "github.com/Syuq/Locket/server/route/api/v2"
	"github.com/Syuq/Locket/store"
)

type TelegramHandler struct {
	store *store.Store
}

func NewTelegramHandler(store *store.Store) *TelegramHandler {
	return &TelegramHandler{store: store}
}

func (t *TelegramHandler) BotToken(ctx context.Context) string {
	if setting, err := t.store.GetWorkspaceSetting(ctx, &store.FindWorkspaceSetting{
		Name: apiv1.SystemSettingTelegramBotTokenName.String(),
	}); err == nil && setting != nil {
		return setting.Value
	}
	return ""
}

const (
	workingMessage = "Working on sending your locket..."
	successMessage = "Success"
)

func (t *TelegramHandler) MessageHandle(ctx context.Context, bot *telegram.Bot, message telegram.Message, attachments []telegram.Attachment) error {
	reply, err := bot.SendReplyMessage(ctx, message.Chat.ID, message.MessageID, workingMessage)
	if err != nil {
		return errors.Wrap(err, "Failed to SendReplyMessage")
	}

	messageSenderID := strconv.FormatInt(message.From.ID, 10)
	var creatorID int32
	userSettingList, err := t.store.ListUserSettings(ctx, &store.FindUserSetting{
		Key: storepb.UserSettingKey_USER_SETTING_TELEGRAM_USER_ID,
	})
	if err != nil {
		return errors.Wrap(err, "Failed to find userSettingList")
	}
	for _, userSetting := range userSettingList {
		if userSetting.GetTelegramUserId() == messageSenderID {
			creatorID = userSetting.UserId
		}
	}

	// If creatorID is not found, ask the user to set the telegram userid in UserSetting of lockets.
	if creatorID == 0 {
		_, err := bot.EditMessage(ctx, message.Chat.ID, reply.MessageID, fmt.Sprintf("Please set your telegram userid %d in UserSetting of lockets", message.From.ID), nil)
		return err
	}

	create := &store.Locket{
		UID:        shortuuid.New(),
		CreatorID:  creatorID,
		Visibility: store.Private,
	}
	if message.Text != nil {
		create.Content = convertToMarkdown(*message.Text, message.Entities)
	}
	if message.Caption != nil {
		create.Content = convertToMarkdown(*message.Caption, message.CaptionEntities)
	}
	if message.ForwardFromChat != nil {
		create.Content += fmt.Sprintf("\n\n[Message link](%s)", message.GetMessageLink())
	}
	locketMessage, err := t.store.CreateLocket(ctx, create)
	if err != nil {
		_, err := bot.EditMessage(ctx, message.Chat.ID, reply.MessageID, fmt.Sprintf("Failed to CreateLocket: %s", err), nil)
		return err
	}

	// Dynamically upsert tags from locket content.
	nodes, err := parser.Parse(tokenizer.Tokenize(create.Content))
	if err != nil {
		return errors.Wrap(err, "Failed to parse content")
	}
	tags := []string{}
	apiv2.TraverseASTNodes(nodes, func(node ast.Node) {
		if tagNode, ok := node.(*ast.Tag); ok {
			tag := tagNode.Content
			if !slices.Contains(tags, tag) {
				tags = append(tags, tag)
			}
		}
	})
	for _, tag := range tags {
		_, err := t.store.UpsertTag(ctx, &store.Tag{
			Name:      tag,
			CreatorID: creatorID,
		})
		if err != nil {
			return errors.Wrap(err, "Failed to upsert tag")
		}
	}

	// Create locket related resources.
	for _, attachment := range attachments {
		// Fill the common field of create
		create := store.Resource{
			UID:       shortuuid.New(),
			CreatorID: creatorID,
			Filename:  filepath.Base(attachment.FileName),
			Type:      attachment.GetMimeType(),
			Size:      attachment.FileSize,
			LocketID:    &locketMessage.ID,
		}

		err := apiv1.SaveResourceBlob(ctx, t.store, &create, bytes.NewReader(attachment.Data))
		if err != nil {
			_, err := bot.EditMessage(ctx, message.Chat.ID, reply.MessageID, fmt.Sprintf("Failed to SaveResourceBlob: %s", err), nil)
			return err
		}

		_, err = t.store.CreateResource(ctx, &create)
		if err != nil {
			_, err := bot.EditMessage(ctx, message.Chat.ID, reply.MessageID, fmt.Sprintf("Failed to CreateResource: %s", err), nil)
			return err
		}
	}

	keyboard := generateKeyboardForLocketID(locketMessage.ID)
	_, err = bot.EditMessage(ctx, message.Chat.ID, reply.MessageID, fmt.Sprintf("Saved as %s Locket %d", locketMessage.Visibility, locketMessage.ID), keyboard)
	_ = t.dispatchLocketRelatedWebhook(ctx, *locketMessage, "lockets.locket.created")
	return err
}

func (t *TelegramHandler) CallbackQueryHandle(ctx context.Context, bot *telegram.Bot, callbackQuery telegram.CallbackQuery) error {
	var locketID int32
	var visibility store.Visibility
	n, err := fmt.Sscanf(callbackQuery.Data, "%s %d", &visibility, &locketID)
	if err != nil || n != 2 {
		return bot.AnswerCallbackQuery(ctx, callbackQuery.ID, fmt.Sprintf("Failed to parse callbackQuery.Data %s", callbackQuery.Data))
	}

	locket, err := t.store.GetLocket(ctx, &store.FindLocket{
		ID: &locketID,
	})
	if err != nil {
		return bot.AnswerCallbackQuery(ctx, callbackQuery.ID, fmt.Sprintf("Failed to call FindLocket %s", err))
	}
	if locket == nil {
		_, err = bot.EditMessage(ctx, callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, fmt.Sprintf("Locket %d not found", locketID), nil)
		if err != nil {
			return bot.AnswerCallbackQuery(ctx, callbackQuery.ID, fmt.Sprintf("Failed to EditMessage %s", err))
		}
		return bot.AnswerCallbackQuery(ctx, callbackQuery.ID, fmt.Sprintf("Locket %d not found, possibly deleted elsewhere", locketID))
	}

	setting, err := t.store.GetWorkspaceSetting(ctx, &store.FindWorkspaceSetting{
		Name: apiv1.SystemSettingDisablePublicLocketsName.String(),
	})
	if err != nil {
		return bot.AnswerCallbackQuery(ctx, callbackQuery.ID, fmt.Sprintf("Failed to get workspace setting %s", err))
	}
	if setting != nil && setting.Value != "" {
		disablePublicLocket := false
		err = json.Unmarshal([]byte(setting.Value), &disablePublicLocket)
		if err != nil {
			return bot.AnswerCallbackQuery(ctx, callbackQuery.ID, fmt.Sprintf("Failed to get workspace setting %s", err))
		}
		if disablePublicLocket && visibility == store.Public {
			return bot.AnswerCallbackQuery(ctx, callbackQuery.ID, fmt.Sprintf("Failed to changing Locket %d to %s\n(workspace disallowed public locket)", locketID, visibility))
		}
	}

	update := store.UpdateLocket{
		ID:         locketID,
		Visibility: &visibility,
	}
	err = t.store.UpdateLocket(ctx, &update)
	if err != nil {
		return bot.AnswerCallbackQuery(ctx, callbackQuery.ID, fmt.Sprintf("Failed to call UpdateLocket %s", err))
	}

	keyboard := generateKeyboardForLocketID(locketID)
	_, err = bot.EditMessage(ctx, callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, fmt.Sprintf("Saved as %s Locket %d", visibility, locketID), keyboard)
	if err != nil {
		return bot.AnswerCallbackQuery(ctx, callbackQuery.ID, fmt.Sprintf("Failed to EditMessage %s", err))
	}

	err = bot.AnswerCallbackQuery(ctx, callbackQuery.ID, fmt.Sprintf("Success changing Locket %d to %s", locketID, visibility))

	locket, webhookErr := t.store.GetLocket(ctx, &store.FindLocket{
		ID: &locketID,
	})
	if webhookErr == nil {
		_ = t.dispatchLocketRelatedWebhook(ctx, *locket, "lockets.locket.updated")
	}
	return err
}

func generateKeyboardForLocketID(id int32) [][]telegram.InlineKeyboardButton {
	allVisibility := []store.Visibility{
		store.Public,
		store.Protected,
		store.Private,
	}

	buttons := make([]telegram.InlineKeyboardButton, 0, len(allVisibility))
	for _, v := range allVisibility {
		button := telegram.InlineKeyboardButton{
			Text:         v.String(),
			CallbackData: fmt.Sprintf("%s %d", v, id),
		}
		buttons = append(buttons, button)
	}

	return [][]telegram.InlineKeyboardButton{buttons}
}

func convertToMarkdown(text string, messageEntities []telegram.MessageEntity) string {
	insertions := make(map[int]string)

	for _, e := range messageEntities {
		var before, after string

		// this is supported by the current markdown
		switch e.Type {
		case telegram.Bold:
			before = "**"
			after = "**"
		case telegram.Italic:
			before = "*"
			after = "*"
		case telegram.Strikethrough:
			before = "~~"
			after = "~~"
		case telegram.Code:
			before = "`"
			after = "`"
		case telegram.Pre:
			before = "```" + e.Language
			after = "```"
		case telegram.TextLink:
			before = "["
			after = fmt.Sprintf(`](%s)`, e.URL)
		case telegram.Spoiler:
			before = "||"
			after = "||"
		}

		if before != "" {
			insertions[e.Offset] += before
			insertions[e.Offset+e.Length] = after + insertions[e.Offset+e.Length]
		}
	}

	input := []rune(text)
	var output []rune
	utf16pos := 0

	for i := 0; i < len(input); i++ {
		output = append(output, []rune(insertions[utf16pos])...)
		output = append(output, input[i])
		utf16pos += len(utf16.Encode([]rune{input[i]}))
	}
	output = append(output, []rune(insertions[utf16pos])...)

	return string(output)
}

func (t *TelegramHandler) dispatchLocketRelatedWebhook(ctx context.Context, locket store.Locket, activityType string) error {
	webhooks, err := t.store.ListWebhooks(ctx, &store.FindWebhook{
		CreatorID: &locket.CreatorID,
	})
	if err != nil {
		return err
	}
	for _, hook := range webhooks {
		payload := t.convertLocketToWebhookPayload(ctx, locket)
		payload.ActivityType = activityType
		payload.URL = hook.Url
		err := webhook.Post(*payload)
		if err != nil {
			return errors.Wrap(err, "failed to post webhook")
		}
	}
	return nil
}

func (t *TelegramHandler) convertLocketToWebhookPayload(ctx context.Context, locket store.Locket) (payload *webhook.WebhookPayload) {
	payload = &webhook.WebhookPayload{
		CreatorID: locket.CreatorID,
		CreatedTs: time.Now().Unix(),
		Locket: &webhook.Locket{
			ID:           locket.ID,
			CreatorID:    locket.CreatorID,
			CreatedTs:    locket.CreatedTs,
			UpdatedTs:    locket.UpdatedTs,
			Content:      locket.Content,
			Visibility:   locket.Visibility.String(),
			Pinned:       locket.Pinned,
			ResourceList: make([]*webhook.Resource, 0),
			RelationList: make([]*webhook.LocketRelation, 0),
		},
	}

	resourceList, err := t.store.ListResources(ctx, &store.FindResource{
		LocketID: &locket.ID,
	})

	if err != nil {
		return payload
	}
	for _, resource := range resourceList {
		payload.Locket.ResourceList = append(payload.Locket.ResourceList, &webhook.Resource{
			ID:           resource.ID,
			CreatorID:    resource.CreatorID,
			CreatedTs:    resource.CreatedTs,
			UpdatedTs:    resource.UpdatedTs,
			Filename:     resource.Filename,
			Type:         resource.Type,
			Size:         resource.Size,
			InternalPath: resource.InternalPath,
			ExternalLink: resource.ExternalLink,
		})
	}

	relationList, err := t.store.ListLocketRelations(ctx, &store.FindLocketRelation{
		LocketID: &locket.ID,
	})

	if err != nil {
		return payload
	}

	for _, relation := range relationList {
		payload.Locket.RelationList = append(payload.Locket.RelationList, &webhook.LocketRelation{
			LocketID:        relation.LocketID,
			RelatedLocketID: relation.RelatedLocketID,
			Type:          string(relation.Type),
		})
	}
	return payload
}
