package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lithammer/shortuuid/v4"
	"github.com/pkg/errors"

	"github.com/Syuq/Locket/internal/util"
	"github.com/Syuq/Locket/plugin/webhook"
	storepb "github.com/Syuq/Locket/proto/gen/store"
	"github.com/Syuq/Locket/store"
)

// Visibility is the type of a visibility.
type Visibility string

const (
	// Public is the PUBLIC visibility.
	Public Visibility = "PUBLIC"
	// Protected is the PROTECTED visibility.
	Protected Visibility = "PROTECTED"
	// Private is the PRIVATE visibility.
	Private Visibility = "PRIVATE"
)

func (v Visibility) String() string {
	switch v {
	case Public:
		return "PUBLIC"
	case Protected:
		return "PROTECTED"
	case Private:
		return "PRIVATE"
	}
	return "PRIVATE"
}

type Locket struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`

	// Standard fields
	RowStatus RowStatus `json:"rowStatus"`
	CreatorID int32     `json:"creatorId"`
	CreatedTs int64     `json:"createdTs"`
	UpdatedTs int64     `json:"updatedTs"`

	// Domain specific fields
	DisplayTs  int64      `json:"displayTs"`
	Content    string     `json:"content"`
	Visibility Visibility `json:"visibility"`
	Pinned     bool       `json:"pinned"`

	// Related fields
	CreatorName     string          `json:"creatorName"`
	CreatorUsername string          `json:"creatorUsername"`
	ResourceList    []*Resource     `json:"resourceList"`
	RelationList    []*LocketRelation `json:"relationList"`
}

type CreateLocketRequest struct {
	// Standard fields
	CreatorID int32  `json:"-"`
	CreatedTs *int64 `json:"createdTs"`

	// Domain specific fields
	Visibility Visibility `json:"visibility"`
	Content    string     `json:"content"`

	// Related fields
	ResourceIDList []int32                      `json:"resourceIdList"`
	RelationList   []*UpsertLocketRelationRequest `json:"relationList"`
}

type PatchLocketRequest struct {
	ID int32 `json:"-"`

	// Standard fields
	CreatedTs *int64 `json:"createdTs"`
	UpdatedTs *int64
	RowStatus *RowStatus `json:"rowStatus"`

	// Domain specific fields
	Content    *string     `json:"content"`
	Visibility *Visibility `json:"visibility"`

	// Related fields
	ResourceIDList []int32                      `json:"resourceIdList"`
	RelationList   []*UpsertLocketRelationRequest `json:"relationList"`
}

type FindLocketRequest struct {
	ID *int32

	// Standard fields
	RowStatus *RowStatus
	CreatorID *int32

	// Domain specific fields
	Pinned         *bool
	ContentSearch  []string
	VisibilityList []Visibility

	// Pagination
	Limit  *int
	Offset *int
}

// maxContentLength means the max locket content bytes is 1MB.
const maxContentLength = 1 << 30

func (s *APIV1Service) registerLocketRoutes(g *echo.Group) {
	g.GET("/locket", s.GetLocketList)
	g.POST("/locket", s.CreateLocket)
	g.GET("/locket/all", s.GetAllLockets)
	g.GET("/locket/stats", s.GetLocketStats)
	g.GET("/locket/:locketId", s.GetLocket)
	g.PATCH("/locket/:locketId", s.UpdateLocket)
	g.DELETE("/locket/:locketId", s.DeleteLocket)
}

// GetLocketList godoc
//
//	@Summary	Get a list of lockets matching optional filters
//	@Tags		locket
//	@Produce	json
//	@Param		creatorId		query		int				false	"Creator ID"
//	@Param		creatorUsername	query		string			false	"Creator username"
//	@Param		rowStatus		query		store.RowStatus	false	"Row status"
//	@Param		pinned			query		bool			false	"Pinned"
//	@Param		tag				query		string			false	"Search for tag. Do not append #"
//	@Param		content			query		string			false	"Search for content"
//	@Param		limit			query		int				false	"Limit"
//	@Param		offset			query		int				false	"Offset"
//	@Success	200				{object}	[]store.Locket	"Locket list"
//	@Failure	400				{object}	nil				"Missing user to find locket"
//	@Failure	500				{object}	nil				"Failed to get locket display with updated ts setting value | Failed to fetch locket list | Failed to compose locket response"
//	@Router		/api/v1/locket [GET]
func (s *APIV1Service) GetLocketList(c echo.Context) error {
	ctx := c.Request().Context()
	find := &store.FindLocket{
		OrderByPinned: true,
	}
	if userID, err := util.ConvertStringToInt32(c.QueryParam("creatorId")); err == nil {
		find.CreatorID = &userID
	}

	if username := c.QueryParam("creatorUsername"); username != "" {
		user, _ := s.Store.GetUser(ctx, &store.FindUser{Username: &username})
		if user != nil {
			find.CreatorID = &user.ID
		}
	}

	currentUserID, ok := c.Get(userIDContextKey).(int32)
	if !ok {
		// Anonymous use should only fetch PUBLIC lockets with specified user
		if find.CreatorID == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing user to find locket")
		}
		find.VisibilityList = []store.Visibility{store.Public}
	} else {
		// Authorized user can fetch all PUBLIC/PROTECTED locket
		visibilityList := []store.Visibility{store.Public, store.Protected}

		// If Creator is authorized user (as default), PRIVATE locket is OK
		if find.CreatorID == nil || *find.CreatorID == currentUserID {
			find.CreatorID = &currentUserID
			visibilityList = append(visibilityList, store.Private)
		}
		find.VisibilityList = visibilityList
	}

	rowStatus := store.RowStatus(c.QueryParam("rowStatus"))
	if rowStatus != "" {
		find.RowStatus = &rowStatus
	}

	contentSearch := []string{}
	tag := c.QueryParam("tag")
	if tag != "" {
		contentSearch = append(contentSearch, "#"+tag)
	}
	content := c.QueryParam("content")
	if content != "" {
		contentSearch = append(contentSearch, content)
	}
	find.ContentSearch = contentSearch

	if limit, err := strconv.Atoi(c.QueryParam("limit")); err == nil {
		find.Limit = &limit
	}
	if offset, err := strconv.Atoi(c.QueryParam("offset")); err == nil {
		find.Offset = &offset
	}

	locketDisplayWithUpdatedTs, err := s.getLocketDisplayWithUpdatedTsSettingValue(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get locket display with updated ts setting value").SetInternal(err)
	}
	if locketDisplayWithUpdatedTs {
		find.OrderByUpdatedTs = true
	}

	list, err := s.Store.ListLockets(ctx, find)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch locket list").SetInternal(err)
	}
	locketResponseList := []*Locket{}
	for _, locket := range list {
		locketResponse, err := s.convertLocketFromStore(ctx, locket)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to compose locket response").SetInternal(err)
		}
		locketResponseList = append(locketResponseList, locketResponse)
	}
	return c.JSON(http.StatusOK, locketResponseList)
}

// CreateLocket godoc
//
//	@Summary		Create a locket
//	@Description	Visibility can be PUBLIC, PROTECTED or PRIVATE
//	@Description	*You should omit fields to use their default values
//	@Tags			locket
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreateLocketRequest	true	"Request object."
//	@Success		200		{object}	store.Locket			"Stored locket"
//	@Failure		400		{object}	nil					"Malformatted post locket request | Content size overflow, up to 1MB"
//	@Failure		401		{object}	nil					"Missing user in session"
//	@Failure		404		{object}	nil					"User not found | Locket not found: %d"
//	@Failure		500		{object}	nil					"Failed to find user setting | Failed to unmarshal user setting value | Failed to find system setting | Failed to unmarshal system setting | Failed to find user | Failed to create locket | Failed to create activity | Failed to upsert locket resource | Failed to upsert locket relation | Failed to compose locket | Failed to compose locket response"
//	@Router			/api/v1/locket [POST]
//
// NOTES:
// - It's currently possible to create phantom resources and relations. Phantom relations will trigger backend 404's when fetching locket.
func (s *APIV1Service) CreateLocket(c echo.Context) error {
	ctx := c.Request().Context()
	userID, ok := c.Get(userIDContextKey).(int32)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Missing user in session")
	}

	createLocketRequest := &CreateLocketRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(createLocketRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Malformatted post locket request").SetInternal(err)
	}
	if len(createLocketRequest.Content) > maxContentLength {
		return echo.NewHTTPError(http.StatusBadRequest, "Content size overflow, up to 1MB")
	}

	if createLocketRequest.Visibility == "" {
		userLocketVisibilitySetting, err := s.Store.GetUserSetting(ctx, &store.FindUserSetting{
			UserID: &userID,
			Key:    storepb.UserSettingKey_USER_SETTING_LOCKET_VISIBILITY,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find user setting").SetInternal(err)
		}
		if userLocketVisibilitySetting != nil {
			createLocketRequest.Visibility = Visibility(userLocketVisibilitySetting.GetLocketVisibility())
		} else {
			// Private is the default locket visibility.
			createLocketRequest.Visibility = Private
		}
	}

	// Find disable public lockets system setting.
	disablePublicLocketsSystemSetting, err := s.Store.GetWorkspaceSetting(ctx, &store.FindWorkspaceSetting{
		Name: SystemSettingDisablePublicLocketsName.String(),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find system setting").SetInternal(err)
	}
	if disablePublicLocketsSystemSetting != nil {
		disablePublicLockets := false
		err = json.Unmarshal([]byte(disablePublicLocketsSystemSetting.Value), &disablePublicLockets)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to unmarshal system setting").SetInternal(err)
		}
		if disablePublicLockets {
			user, err := s.Store.GetUser(ctx, &store.FindUser{
				ID: &userID,
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find user").SetInternal(err)
			}
			if user == nil {
				return echo.NewHTTPError(http.StatusNotFound, "User not found")
			}
			// Enforce normal user to create private locket if public lockets are disabled.
			if user.Role == store.RoleUser {
				createLocketRequest.Visibility = Private
			}
		}
	}

	createLocketRequest.CreatorID = userID
	locket, err := s.Store.CreateLocket(ctx, convertCreateLocketRequestToLocketMessage(createLocketRequest))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create locket").SetInternal(err)
	}

	for _, resourceID := range createLocketRequest.ResourceIDList {
		if _, err := s.Store.UpdateResource(ctx, &store.UpdateResource{
			ID:     resourceID,
			LocketID: &locket.ID,
		}); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to upsert locket resource").SetInternal(err)
		}
	}

	for _, locketRelationUpsert := range createLocketRequest.RelationList {
		if _, err := s.Store.UpsertLocketRelation(ctx, &store.LocketRelation{
			LocketID:        locket.ID,
			RelatedLocketID: locketRelationUpsert.RelatedLocketID,
			Type:          store.LocketRelationType(locketRelationUpsert.Type),
		}); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to upsert locket relation").SetInternal(err)
		}
		if locket.Visibility != store.Private && locketRelationUpsert.Type == LocketRelationComment {
			relatedLocket, err := s.Store.GetLocket(ctx, &store.FindLocket{
				ID: &locketRelationUpsert.RelatedLocketID,
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get related locket").SetInternal(err)
			}
			if relatedLocket.CreatorID != locket.CreatorID {
				activity, err := s.Store.CreateActivity(ctx, &store.Activity{
					CreatorID: locket.CreatorID,
					Type:      store.ActivityTypeLocketComment,
					Level:     store.ActivityLevelInfo,
					Payload: &storepb.ActivityPayload{
						LocketComment: &storepb.ActivityLocketCommentPayload{
							LocketId:        locket.ID,
							RelatedLocketId: locketRelationUpsert.RelatedLocketID,
						},
					},
				})
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create activity").SetInternal(err)
				}
				if _, err := s.Store.CreateInbox(ctx, &store.Inbox{
					SenderID:   locket.CreatorID,
					ReceiverID: relatedLocket.CreatorID,
					Status:     store.UNREAD,
					Message: &storepb.InboxMessage{
						Type:       storepb.InboxMessage_TYPE_LOCKET_COMMENT,
						ActivityId: &activity.ID,
					},
				}); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create inbox").SetInternal(err)
				}
			}
		}
	}

	composedLocket, err := s.Store.GetLocket(ctx, &store.FindLocket{
		ID: &locket.ID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to compose locket").SetInternal(err)
	}
	if composedLocket == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Locket not found: %d", locket.ID))
	}

	locketResponse, err := s.convertLocketFromStore(ctx, composedLocket)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to compose locket response").SetInternal(err)
	}

	// Send notification to telegram if locket is not private.
	if locketResponse.Visibility != Private {
		// fetch all telegram UserID
		userSettings, err := s.Store.ListUserSettings(ctx, &store.FindUserSetting{Key: storepb.UserSettingKey_USER_SETTING_TELEGRAM_USER_ID})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to ListUserSettings").SetInternal(err)
		}
		for _, userSetting := range userSettings {
			tgUserID, err := strconv.ParseInt(userSetting.GetTelegramUserId(), 10, 64)
			if err != nil {
				continue
			}

			// send notification to telegram
			content := locketResponse.CreatorName + " Says:\n\n" + locketResponse.Content
			_, err = s.telegramBot.SendMessage(ctx, tgUserID, content)
			if err != nil {
				continue
			}
		}
	}
	// Try to dispatch webhook when locket is created.
	if err := s.DispatchLocketCreatedWebhook(ctx, locketResponse); err != nil {
		slog.Warn("Failed to dispatch locket created webhook", err)
	}

	return c.JSON(http.StatusOK, locketResponse)
}

// GetAllLockets godoc
//
//	@Summary		Get a list of public lockets matching optional filters
//	@Description	This should also list protected lockets if the user is logged in
//	@Description	Authentication is optional
//	@Tags			locket
//	@Produce		json
//	@Param			limit	query		int				false	"Limit"
//	@Param			offset	query		int				false	"Offset"
//	@Success		200		{object}	[]store.Locket	"Locket list"
//	@Failure		500		{object}	nil				"Failed to get locket display with updated ts setting value | Failed to fetch all locket list | Failed to compose locket response"
//	@Router			/api/v1/locket/all [GET]
//
//	NOTES:
//	- creatorUsername is listed at ./web/src/helpers/api.ts:82, but it's not present here
func (s *APIV1Service) GetAllLockets(c echo.Context) error {
	ctx := c.Request().Context()
	locketFind := &store.FindLocket{}
	_, ok := c.Get(userIDContextKey).(int32)
	if !ok {
		locketFind.VisibilityList = []store.Visibility{store.Public}
	} else {
		locketFind.VisibilityList = []store.Visibility{store.Public, store.Protected}
	}

	if limit, err := strconv.Atoi(c.QueryParam("limit")); err == nil {
		locketFind.Limit = &limit
	}
	if offset, err := strconv.Atoi(c.QueryParam("offset")); err == nil {
		locketFind.Offset = &offset
	}

	// Only fetch normal status lockets.
	normalStatus := store.Normal
	locketFind.RowStatus = &normalStatus

	locketDisplayWithUpdatedTs, err := s.getLocketDisplayWithUpdatedTsSettingValue(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get locket display with updated ts setting value").SetInternal(err)
	}
	if locketDisplayWithUpdatedTs {
		locketFind.OrderByUpdatedTs = true
	}

	list, err := s.Store.ListLockets(ctx, locketFind)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch all locket list").SetInternal(err)
	}
	locketResponseList := []*Locket{}
	for _, locket := range list {
		locketResponse, err := s.convertLocketFromStore(ctx, locket)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to compose locket response").SetInternal(err)
		}
		locketResponseList = append(locketResponseList, locketResponse)
	}
	return c.JSON(http.StatusOK, locketResponseList)
}

// GetLocketStats godoc
//
//	@Summary		Get locket stats by creator ID or username
//	@Description	Used to generate the heatmap
//	@Tags			locket
//	@Produce		json
//	@Param			creatorId		query		int		false	"Creator ID"
//	@Param			creatorUsername	query		string	false	"Creator username"
//	@Success		200				{object}	[]int	"Locket createdTs list"
//	@Failure		400				{object}	nil		"Missing user id to find locket"
//	@Failure		500				{object}	nil		"Failed to get locket display with updated ts setting value | Failed to find locket list | Failed to compose locket response"
//	@Router			/api/v1/locket/stats [GET]
func (s *APIV1Service) GetLocketStats(c echo.Context) error {
	ctx := c.Request().Context()
	normalStatus := store.Normal
	findLocketMessage := &store.FindLocket{
		RowStatus:      &normalStatus,
		ExcludeContent: true,
	}
	if creatorID, err := util.ConvertStringToInt32(c.QueryParam("creatorId")); err == nil {
		findLocketMessage.CreatorID = &creatorID
	}

	if username := c.QueryParam("creatorUsername"); username != "" {
		user, _ := s.Store.GetUser(ctx, &store.FindUser{Username: &username})
		if user != nil {
			findLocketMessage.CreatorID = &user.ID
		}
	}

	if findLocketMessage.CreatorID == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing user id to find locket")
	}

	currentUserID, ok := c.Get(userIDContextKey).(int32)
	if !ok {
		findLocketMessage.VisibilityList = []store.Visibility{store.Public}
	} else {
		if *findLocketMessage.CreatorID != currentUserID {
			findLocketMessage.VisibilityList = []store.Visibility{store.Public, store.Protected}
		} else {
			findLocketMessage.VisibilityList = []store.Visibility{store.Public, store.Protected, store.Private}
		}
	}

	locketDisplayWithUpdatedTs, err := s.getLocketDisplayWithUpdatedTsSettingValue(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get locket display with updated ts setting value").SetInternal(err)
	}
	if locketDisplayWithUpdatedTs {
		findLocketMessage.OrderByUpdatedTs = true
	}

	list, err := s.Store.ListLockets(ctx, findLocketMessage)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find locket list").SetInternal(err)
	}

	displayTsList := []int64{}
	if locketDisplayWithUpdatedTs {
		for _, locket := range list {
			displayTsList = append(displayTsList, locket.UpdatedTs)
		}
	} else {
		for _, locket := range list {
			displayTsList = append(displayTsList, locket.CreatedTs)
		}
	}
	return c.JSON(http.StatusOK, displayTsList)
}

// GetLocket godoc
//
//	@Summary	Get locket by ID
//	@Tags		locket
//	@Produce	json
//	@Param		locketId	path		int				true	"Locket ID"
//	@Success	200		{object}	[]store.Locket	"Locket list"
//	@Failure	400		{object}	nil				"ID is not a number: %s"
//	@Failure	401		{object}	nil				"Missing user in session"
//	@Failure	403		{object}	nil				"this locket is private only | this locket is protected, missing user in session
//	@Failure	404		{object}	nil				"Locket not found: %d"
//	@Failure	500		{object}	nil				"Failed to find locket by ID: %v | Failed to compose locket response"
//	@Router		/api/v1/locket/{locketId} [GET]
func (s *APIV1Service) GetLocket(c echo.Context) error {
	ctx := c.Request().Context()
	locketID, err := util.ConvertStringToInt32(c.Param("locketId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("locketId"))).SetInternal(err)
	}

	locket, err := s.Store.GetLocket(ctx, &store.FindLocket{
		ID: &locketID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to find locket by ID: %v", locketID)).SetInternal(err)
	}
	if locket == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Locket not found: %d", locketID))
	}

	userID, ok := c.Get(userIDContextKey).(int32)
	if locket.Visibility == store.Private {
		if !ok || locket.CreatorID != userID {
			return echo.NewHTTPError(http.StatusForbidden, "this locket is private only")
		}
	} else if locket.Visibility == store.Protected {
		if !ok {
			return echo.NewHTTPError(http.StatusForbidden, "this locket is protected, missing user in session")
		}
	}
	locketResponse, err := s.convertLocketFromStore(ctx, locket)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to compose locket response").SetInternal(err)
	}
	return c.JSON(http.StatusOK, locketResponse)
}

// DeleteLocket godoc
//
//	@Summary	Delete locket by ID
//	@Tags		locket
//	@Produce	json
//	@Param		locketId	path		int		true	"Locket ID to delete"
//	@Success	200		{boolean}	true	"Locket deleted"
//	@Failure	400		{object}	nil		"ID is not a number: %s"
//	@Failure	401		{object}	nil		"Missing user in session | Unauthorized"
//	@Failure	404		{object}	nil		"Locket not found: %d"
//	@Failure	500		{object}	nil		"Failed to find locket | Failed to delete locket ID: %v"
//	@Router		/api/v1/locket/{locketId} [DELETE]
func (s *APIV1Service) DeleteLocket(c echo.Context) error {
	ctx := c.Request().Context()
	userID, ok := c.Get(userIDContextKey).(int32)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Missing user in session")
	}
	locketID, err := util.ConvertStringToInt32(c.Param("locketId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("locketId"))).SetInternal(err)
	}

	locket, err := s.Store.GetLocket(ctx, &store.FindLocket{
		ID: &locketID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find locket").SetInternal(err)
	}
	if locket == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Locket not found: %d", locketID))
	}
	if locket.CreatorID != userID {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	if locketMessage, err := s.convertLocketFromStore(ctx, locket); err == nil {
		// Try to dispatch webhook when locket is deleted.
		if err := s.DispatchLocketDeletedWebhook(ctx, locketMessage); err != nil {
			slog.Warn("Failed to dispatch locket deleted webhook", err)
		}
	}

	if err := s.Store.DeleteLocket(ctx, &store.DeleteLocket{
		ID: locketID,
	}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to delete locket ID: %v", locketID)).SetInternal(err)
	}
	return c.JSON(http.StatusOK, true)
}

// UpdateLocket godoc
//
//	@Summary		Update a locket
//	@Description	Visibility can be PUBLIC, PROTECTED or PRIVATE
//	@Description	*You should omit fields to use their default values
//	@Tags			locket
//	@Accept			json
//	@Produce		json
//	@Param			locketId	path		int					true	"ID of locket to update"
//	@Param			body	body		PatchLocketRequest	true	"Patched object."
//	@Success		200		{object}	store.Locket			"Stored locket"
//	@Failure		400		{object}	nil					"ID is not a number: %s | Malformatted patch locket request | Content size overflow, up to 1MB"
//	@Failure		401		{object}	nil					"Missing user in session | Unauthorized"
//	@Failure		404		{object}	nil					"Locket not found: %d"
//	@Failure		500		{object}	nil					"Failed to find locket | Failed to patch locket | Failed to upsert locket resource | Failed to delete locket resource | Failed to compose locket response"
//	@Router			/api/v1/locket/{locketId} [PATCH]
//
// NOTES:
// - It's currently possible to create phantom resources and relations. Phantom relations will trigger backend 404's when fetching locket.
// - Passing 0 to createdTs and updatedTs will set them to 0 in the database, which is probably unwanted.
func (s *APIV1Service) UpdateLocket(c echo.Context) error {
	ctx := c.Request().Context()
	userID, ok := c.Get(userIDContextKey).(int32)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Missing user in session")
	}

	locketID, err := util.ConvertStringToInt32(c.Param("locketId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("locketId"))).SetInternal(err)
	}

	locket, err := s.Store.GetLocket(ctx, &store.FindLocket{
		ID: &locketID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find locket").SetInternal(err)
	}
	if locket == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Locket not found: %d", locketID))
	}
	if locket.CreatorID != userID {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	currentTs := time.Now().Unix()
	patchLocketRequest := &PatchLocketRequest{
		ID:        locketID,
		UpdatedTs: &currentTs,
	}
	if err := json.NewDecoder(c.Request().Body).Decode(patchLocketRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Malformatted patch locket request").SetInternal(err)
	}

	if patchLocketRequest.Content != nil && len(*patchLocketRequest.Content) > maxContentLength {
		return echo.NewHTTPError(http.StatusBadRequest, "Content size overflow, up to 1MB").SetInternal(err)
	}

	updateLocketMessage := &store.UpdateLocket{
		ID:        locketID,
		CreatedTs: patchLocketRequest.CreatedTs,
		UpdatedTs: patchLocketRequest.UpdatedTs,
		Content:   patchLocketRequest.Content,
	}
	if patchLocketRequest.RowStatus != nil {
		rowStatus := store.RowStatus(patchLocketRequest.RowStatus.String())
		updateLocketMessage.RowStatus = &rowStatus
	}
	if patchLocketRequest.Visibility != nil {
		visibility := store.Visibility(patchLocketRequest.Visibility.String())
		updateLocketMessage.Visibility = &visibility
		// Find disable public lockets system setting.
		disablePublicLocketsSystemSetting, err := s.Store.GetWorkspaceSetting(ctx, &store.FindWorkspaceSetting{
			Name: SystemSettingDisablePublicLocketsName.String(),
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find system setting").SetInternal(err)
		}
		if disablePublicLocketsSystemSetting != nil {
			disablePublicLockets := false
			err = json.Unmarshal([]byte(disablePublicLocketsSystemSetting.Value), &disablePublicLockets)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to unmarshal system setting").SetInternal(err)
			}
			if disablePublicLockets {
				user, err := s.Store.GetUser(ctx, &store.FindUser{
					ID: &userID,
				})
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find user").SetInternal(err)
				}
				if user == nil {
					return echo.NewHTTPError(http.StatusNotFound, "User not found")
				}
				// Enforce normal user to save as private locket if public lockets are disabled.
				if user.Role == store.RoleUser {
					visibility = store.Visibility("PRIVATE")
					updateLocketMessage.Visibility = &visibility
				}
			}
		}
	}

	err = s.Store.UpdateLocket(ctx, updateLocketMessage)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to patch locket").SetInternal(err)
	}
	locket, err = s.Store.GetLocket(ctx, &store.FindLocket{ID: &locketID})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find locket").SetInternal(err)
	}
	if locket == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Locket not found: %d", locketID))
	}

	locketMessage, err := s.convertLocketFromStore(ctx, locket)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to compose locket").SetInternal(err)
	}
	if patchLocketRequest.ResourceIDList != nil {
		originResourceIDList := []int32{}
		for _, resource := range locketMessage.ResourceList {
			originResourceIDList = append(originResourceIDList, resource.ID)
		}
		addedResourceIDList, removedResourceIDList := getIDListDiff(originResourceIDList, patchLocketRequest.ResourceIDList)
		for _, resourceID := range addedResourceIDList {
			if _, err := s.Store.UpdateResource(ctx, &store.UpdateResource{
				ID:     resourceID,
				LocketID: &locket.ID,
			}); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to upsert locket resource").SetInternal(err)
			}
		}
		for _, resourceID := range removedResourceIDList {
			if err := s.Store.DeleteResource(ctx, &store.DeleteResource{
				ID: resourceID,
			}); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete resource").SetInternal(err)
			}
		}
	}

	if patchLocketRequest.RelationList != nil {
		patchLocketRelationList := make([]*LocketRelation, 0)
		for _, locketRelation := range patchLocketRequest.RelationList {
			patchLocketRelationList = append(patchLocketRelationList, &LocketRelation{
				LocketID:        locket.ID,
				RelatedLocketID: locketRelation.RelatedLocketID,
				Type:          locketRelation.Type,
			})
		}
		addedLocketRelationList, removedLocketRelationList := getLocketRelationListDiff(locketMessage.RelationList, patchLocketRelationList)
		for _, locketRelation := range addedLocketRelationList {
			if _, err := s.Store.UpsertLocketRelation(ctx, locketRelation); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to upsert locket relation").SetInternal(err)
			}
		}
		for _, locketRelation := range removedLocketRelationList {
			if err := s.Store.DeleteLocketRelation(ctx, &store.DeleteLocketRelation{
				LocketID:        &locket.ID,
				RelatedLocketID: &locketRelation.RelatedLocketID,
				Type:          &locketRelation.Type,
			}); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete locket relation").SetInternal(err)
			}
		}
	}

	locket, err = s.Store.GetLocket(ctx, &store.FindLocket{ID: &locketID})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find locket").SetInternal(err)
	}
	if locket == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Locket not found: %d", locketID))
	}

	locketResponse, err := s.convertLocketFromStore(ctx, locket)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to compose locket response").SetInternal(err)
	}
	// Try to dispatch webhook when locket is updated.
	if err := s.DispatchLocketUpdatedWebhook(ctx, locketResponse); err != nil {
		slog.Error("Failed to dispatch locket updated webhook", err)
	}

	return c.JSON(http.StatusOK, locketResponse)
}

func (s *APIV1Service) convertLocketFromStore(ctx context.Context, locket *store.Locket) (*Locket, error) {
	locketMessage := &Locket{
		ID:         locket.ID,
		Name:       locket.UID,
		RowStatus:  RowStatus(locket.RowStatus.String()),
		CreatorID:  locket.CreatorID,
		CreatedTs:  locket.CreatedTs,
		UpdatedTs:  locket.UpdatedTs,
		Content:    locket.Content,
		Visibility: Visibility(locket.Visibility.String()),
		Pinned:     locket.Pinned,
	}

	// Compose creator name.
	user, err := s.Store.GetUser(ctx, &store.FindUser{
		ID: &locketMessage.CreatorID,
	})
	if err != nil {
		return nil, err
	}
	if user.Nickname != "" {
		locketMessage.CreatorName = user.Nickname
	} else {
		locketMessage.CreatorName = user.Username
	}
	locketMessage.CreatorUsername = user.Username

	// Compose display ts.
	locketMessage.DisplayTs = locketMessage.CreatedTs
	// Find locket display with updated ts setting.
	locketDisplayWithUpdatedTs, err := s.getLocketDisplayWithUpdatedTsSettingValue(ctx)
	if err != nil {
		return nil, err
	}
	if locketDisplayWithUpdatedTs {
		locketMessage.DisplayTs = locketMessage.UpdatedTs
	}

	// Compose related resources.
	resourceList, err := s.Store.ListResources(ctx, &store.FindResource{
		LocketID: &locket.ID,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list resources")
	}
	locketMessage.ResourceList = []*Resource{}
	for _, resource := range resourceList {
		locketMessage.ResourceList = append(locketMessage.ResourceList, convertResourceFromStore(resource))
	}

	// Compose related locket relations.
	relationList := []*LocketRelation{}
	tempList, err := s.Store.ListLocketRelations(ctx, &store.FindLocketRelation{
		LocketID: &locket.ID,
	})
	if err != nil {
		return nil, err
	}
	for _, relation := range tempList {
		relationList = append(relationList, convertLocketRelationFromStore(relation))
	}
	tempList, err = s.Store.ListLocketRelations(ctx, &store.FindLocketRelation{
		RelatedLocketID: &locket.ID,
	})
	if err != nil {
		return nil, err
	}
	for _, relation := range tempList {
		relationList = append(relationList, convertLocketRelationFromStore(relation))
	}
	locketMessage.RelationList = relationList
	return locketMessage, nil
}

func (s *APIV1Service) getLocketDisplayWithUpdatedTsSettingValue(ctx context.Context) (bool, error) {
	locketDisplayWithUpdatedTsSetting, err := s.Store.GetWorkspaceSetting(ctx, &store.FindWorkspaceSetting{
		Name: SystemSettingLocketDisplayWithUpdatedTsName.String(),
	})
	if err != nil {
		return false, errors.Wrap(err, "failed to find system setting")
	}
	locketDisplayWithUpdatedTs := false
	if locketDisplayWithUpdatedTsSetting != nil {
		err = json.Unmarshal([]byte(locketDisplayWithUpdatedTsSetting.Value), &locketDisplayWithUpdatedTs)
		if err != nil {
			return false, errors.Wrap(err, "failed to unmarshal system setting value")
		}
	}
	return locketDisplayWithUpdatedTs, nil
}

func convertCreateLocketRequestToLocketMessage(locketCreate *CreateLocketRequest) *store.Locket {
	createdTs := time.Now().Unix()
	if locketCreate.CreatedTs != nil {
		createdTs = *locketCreate.CreatedTs
	}
	return &store.Locket{
		UID:        shortuuid.New(),
		CreatorID:  locketCreate.CreatorID,
		CreatedTs:  createdTs,
		Content:    locketCreate.Content,
		Visibility: store.Visibility(locketCreate.Visibility),
	}
}

func getLocketRelationListDiff(oldList, newList []*LocketRelation) (addedList, removedList []*store.LocketRelation) {
	oldMap := map[string]bool{}
	for _, relation := range oldList {
		oldMap[fmt.Sprintf("%d-%s", relation.RelatedLocketID, relation.Type)] = true
	}
	newMap := map[string]bool{}
	for _, relation := range newList {
		newMap[fmt.Sprintf("%d-%s", relation.RelatedLocketID, relation.Type)] = true
	}
	for _, relation := range oldList {
		key := fmt.Sprintf("%d-%s", relation.RelatedLocketID, relation.Type)
		if !newMap[key] {
			removedList = append(removedList, &store.LocketRelation{
				LocketID:        relation.LocketID,
				RelatedLocketID: relation.RelatedLocketID,
				Type:          store.LocketRelationType(relation.Type),
			})
		}
	}
	for _, relation := range newList {
		key := fmt.Sprintf("%d-%s", relation.RelatedLocketID, relation.Type)
		if !oldMap[key] {
			addedList = append(addedList, &store.LocketRelation{
				LocketID:        relation.LocketID,
				RelatedLocketID: relation.RelatedLocketID,
				Type:          store.LocketRelationType(relation.Type),
			})
		}
	}
	return addedList, removedList
}

func getIDListDiff(oldList, newList []int32) (addedList, removedList []int32) {
	oldMap := map[int32]bool{}
	for _, id := range oldList {
		oldMap[id] = true
	}
	newMap := map[int32]bool{}
	for _, id := range newList {
		newMap[id] = true
	}
	for id := range oldMap {
		if !newMap[id] {
			removedList = append(removedList, id)
		}
	}
	for id := range newMap {
		if !oldMap[id] {
			addedList = append(addedList, id)
		}
	}
	return addedList, removedList
}

// DispatchLocketCreatedWebhook dispatches webhook when locket is created.
func (s *APIV1Service) DispatchLocketCreatedWebhook(ctx context.Context, locket *Locket) error {
	return s.dispatchLocketRelatedWebhook(ctx, locket, "lockets.locket.created")
}

// DispatchLocketUpdatedWebhook dispatches webhook when locket is updated.
func (s *APIV1Service) DispatchLocketUpdatedWebhook(ctx context.Context, locket *Locket) error {
	return s.dispatchLocketRelatedWebhook(ctx, locket, "lockets.locket.updated")
}

// DispatchLocketDeletedWebhook dispatches webhook when locket is deletedd.
func (s *APIV1Service) DispatchLocketDeletedWebhook(ctx context.Context, locket *Locket) error {
	return s.dispatchLocketRelatedWebhook(ctx, locket, "lockets.locket.deleted")
}

func (s *APIV1Service) dispatchLocketRelatedWebhook(ctx context.Context, locket *Locket, activityType string) error {
	webhooks, err := s.Store.ListWebhooks(ctx, &store.FindWebhook{
		CreatorID: &locket.CreatorID,
	})
	if err != nil {
		return err
	}
	for _, hook := range webhooks {
		payload := convertLocketToWebhookPayload(locket)
		payload.ActivityType = activityType
		payload.URL = hook.Url
		err := webhook.Post(*payload)
		if err != nil {
			return errors.Wrap(err, "failed to post webhook")
		}
	}
	return nil
}

func convertLocketToWebhookPayload(locket *Locket) *webhook.WebhookPayload {
	return &webhook.WebhookPayload{
		CreatorID: locket.CreatorID,
		CreatedTs: time.Now().Unix(),
		Locket: &webhook.Locket{
			ID:         locket.ID,
			CreatorID:  locket.CreatorID,
			CreatedTs:  locket.CreatedTs,
			UpdatedTs:  locket.UpdatedTs,
			Content:    locket.Content,
			Visibility: locket.Visibility.String(),
			Pinned:     locket.Pinned,
			ResourceList: func() []*webhook.Resource {
				resources := []*webhook.Resource{}
				for _, resource := range locket.ResourceList {
					resources = append(resources, &webhook.Resource{
						ID:           resource.ID,
						CreatorID:    resource.CreatorID,
						CreatedTs:    resource.CreatedTs,
						UpdatedTs:    resource.UpdatedTs,
						Filename:     resource.Filename,
						InternalPath: resource.InternalPath,
						ExternalLink: resource.ExternalLink,
						Type:         resource.Type,
						Size:         resource.Size,
					})
				}
				return resources
			}(),
			RelationList: func() []*webhook.LocketRelation {
				relations := []*webhook.LocketRelation{}
				for _, relation := range locket.RelationList {
					relations = append(relations, &webhook.LocketRelation{
						LocketID:        relation.LocketID,
						RelatedLocketID: relation.RelatedLocketID,
						Type:          relation.Type.String(),
					})
				}
				return relations
			}(),
		},
	}
}
