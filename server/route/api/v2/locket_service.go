package v2

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/cel-go/cel"
	"github.com/lithammer/shortuuid/v4"
	"github.com/pkg/errors"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Syuq/Locket/internal/util"
	"github.com/Syuq/Locket/plugin/webhook"
	apiv2pb "github.com/Syuq/Locket/proto/gen/api/v2"
	storepb "github.com/Syuq/Locket/proto/gen/store"
	apiv1 "github.com/Syuq/Locket/server/route/api/v1"
	"github.com/Syuq/Locket/store"
)

const (
	DefaultPageSize  = 10
	MaxContentLength = 8 * 1024
	ChunkSize        = 64 * 1024 // 64 KiB
)

func (s *APIV2Service) CreateLocket(ctx context.Context, request *apiv2pb.CreateLocketRequest) (*apiv2pb.CreateLocketResponse, error) {
	user, err := getCurrentUser(ctx, s.Store)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user")
	}
	if user == nil {
		return nil, status.Errorf(codes.PermissionDenied, "permission denied")
	}
	if len(request.Content) > MaxContentLength {
		return nil, status.Errorf(codes.InvalidArgument, "content too long")
	}

	create := &store.Locket{
		UID:        shortuuid.New(),
		CreatorID:  user.ID,
		Content:    request.Content,
		Visibility: convertVisibilityToStore(request.Visibility),
	}
	// Find disable public lockets system setting.
	disablePublicLocketsSystem, err := s.getDisablePublicLocketsSystemSettingValue(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get system setting")
	}
	if disablePublicLocketsSystem && create.Visibility == store.Public {
		return nil, status.Errorf(codes.PermissionDenied, "disable public lockets system setting is enabled")
	}

	locket, err := s.Store.CreateLocket(ctx, create)
	if err != nil {
		return nil, err
	}

	locketMessage, err := s.convertLocketFromStore(ctx, locket)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert locket")
	}
	// Try to dispatch webhook when locket is created.
	if err := s.DispatchLocketCreatedWebhook(ctx, locketMessage); err != nil {
		slog.Warn("Failed to dispatch locket created webhook", err)
	}

	response := &apiv2pb.CreateLocketResponse{
		Locket: locketMessage,
	}
	return response, nil
}

func (s *APIV2Service) ListLockets(ctx context.Context, request *apiv2pb.ListLocketsRequest) (*apiv2pb.ListLocketsResponse, error) {
	locketFind := &store.FindLocket{
		// Exclude comments by default.
		ExcludeComments: true,
	}
	if err := s.buildLocketFindWithFilter(ctx, locketFind, request.Filter); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to build find lockets with filter")
	}

	var limit, offset int
	if request.PageToken != "" {
		var pageToken apiv2pb.PageToken
		if err := unmarshalPageToken(request.PageToken, &pageToken); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid page token: %v", err)
		}
		limit = int(pageToken.Limit)
		offset = int(pageToken.Offset)
	} else {
		limit = int(request.PageSize)
	}
	if limit <= 0 {
		limit = DefaultPageSize
	}
	limitPlusOne := limit + 1
	locketFind.Limit = &limitPlusOne
	locketFind.Offset = &offset
	lockets, err := s.Store.ListLockets(ctx, locketFind)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list lockets: %v", err)
	}

	locketMessages := []*apiv2pb.Locket{}
	nextPageToken := ""
	if len(lockets) == limitPlusOne {
		lockets = lockets[:limit]
		nextPageToken, err = getPageToken(limit, offset+limit)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get next page token, error: %v", err)
		}
	}
	for _, locket := range lockets {
		locketMessage, err := s.convertLocketFromStore(ctx, locket)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert locket")
		}
		locketMessages = append(locketMessages, locketMessage)
	}

	response := &apiv2pb.ListLocketsResponse{
		Lockets:         locketMessages,
		NextPageToken: nextPageToken,
	}
	return response, nil
}

func (s *APIV2Service) SearchLockets(ctx context.Context, request *apiv2pb.SearchLocketsRequest) (*apiv2pb.SearchLocketsResponse, error) {
	defaultSearchLimit := 10
	locketFind := &store.FindLocket{
		// Exclude comments by default.
		ExcludeComments: true,
		Limit:           &defaultSearchLimit,
	}
	err := s.buildLocketFindWithFilter(ctx, locketFind, request.Filter)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to build find lockets with filter")
	}

	lockets, err := s.Store.ListLockets(ctx, locketFind)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to search lockets")
	}

	locketMessages := []*apiv2pb.Locket{}
	for _, locket := range lockets {
		locketMessage, err := s.convertLocketFromStore(ctx, locket)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert locket")
		}
		locketMessages = append(locketMessages, locketMessage)
	}

	response := &apiv2pb.SearchLocketsResponse{
		Lockets: locketMessages,
	}
	return response, nil
}

func (s *APIV2Service) GetLocket(ctx context.Context, request *apiv2pb.GetLocketRequest) (*apiv2pb.GetLocketResponse, error) {
	id, err := ExtractLocketIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid locket name: %v", err)
	}
	locket, err := s.Store.GetLocket(ctx, &store.FindLocket{
		ID: &id,
	})
	if err != nil {
		return nil, err
	}
	if locket == nil {
		return nil, status.Errorf(codes.NotFound, "locket not found")
	}
	if locket.Visibility != store.Public {
		user, err := getCurrentUser(ctx, s.Store)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get user")
		}
		if user == nil {
			return nil, status.Errorf(codes.PermissionDenied, "permission denied")
		}
		if locket.Visibility == store.Private && locket.CreatorID != user.ID {
			return nil, status.Errorf(codes.PermissionDenied, "permission denied")
		}
	}

	locketMessage, err := s.convertLocketFromStore(ctx, locket)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert locket")
	}
	response := &apiv2pb.GetLocketResponse{
		Locket: locketMessage,
	}
	return response, nil
}

func (s *APIV2Service) UpdateLocket(ctx context.Context, request *apiv2pb.UpdateLocketRequest) (*apiv2pb.UpdateLocketResponse, error) {
	id, err := ExtractLocketIDFromName(request.Locket.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid locket name: %v", err)
	}
	if request.UpdateMask == nil || len(request.UpdateMask.Paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update mask is required")
	}

	locket, err := s.Store.GetLocket(ctx, &store.FindLocket{ID: &id})
	if err != nil {
		return nil, err
	}
	if locket == nil {
		return nil, status.Errorf(codes.NotFound, "locket not found")
	}

	user, _ := getCurrentUser(ctx, s.Store)
	if locket.CreatorID != user.ID {
		return nil, status.Errorf(codes.PermissionDenied, "permission denied")
	}

	currentTs := time.Now().Unix()
	update := &store.UpdateLocket{
		ID:        id,
		UpdatedTs: &currentTs,
	}
	for _, path := range request.UpdateMask.Paths {
		if path == "content" {
			update.Content = &request.Locket.Content
		} else if path == "uid" {
			update.UID = &request.Locket.Name
			if !util.UIDMatcher.MatchString(*update.UID) {
				return nil, status.Errorf(codes.InvalidArgument, "invalid resource name")
			}
		} else if path == "visibility" {
			visibility := convertVisibilityToStore(request.Locket.Visibility)
			// Find disable public lockets system setting.
			disablePublicLocketsSystem, err := s.getDisablePublicLocketsSystemSettingValue(ctx)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to get system setting")
			}
			if disablePublicLocketsSystem && visibility == store.Public {
				return nil, status.Errorf(codes.PermissionDenied, "disable public lockets system setting is enabled")
			}
			update.Visibility = &visibility
		} else if path == "row_status" {
			rowStatus := convertRowStatusToStore(request.Locket.RowStatus)
			update.RowStatus = &rowStatus
		} else if path == "created_ts" {
			createdTs := request.Locket.CreateTime.AsTime().Unix()
			update.CreatedTs = &createdTs
		} else if path == "pinned" {
			if _, err := s.Store.UpsertLocketOrganizer(ctx, &store.LocketOrganizer{
				LocketID: id,
				UserID: user.ID,
				Pinned: request.Locket.Pinned,
			}); err != nil {
				return nil, status.Errorf(codes.Internal, "failed to upsert locket organizer")
			}
		}
	}
	if update.Content != nil && len(*update.Content) > MaxContentLength {
		return nil, status.Errorf(codes.InvalidArgument, "content too long")
	}

	if err = s.Store.UpdateLocket(ctx, update); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update locket")
	}

	locket, err = s.Store.GetLocket(ctx, &store.FindLocket{
		ID: &id,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get locket")
	}
	locketMessage, err := s.convertLocketFromStore(ctx, locket)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert locket")
	}
	// Try to dispatch webhook when locket is updated.
	if err := s.DispatchLocketUpdatedWebhook(ctx, locketMessage); err != nil {
		slog.Warn("Failed to dispatch locket updated webhook", err)
	}

	return &apiv2pb.UpdateLocketResponse{
		Locket: locketMessage,
	}, nil
}

func (s *APIV2Service) DeleteLocket(ctx context.Context, request *apiv2pb.DeleteLocketRequest) (*apiv2pb.DeleteLocketResponse, error) {
	id, err := ExtractLocketIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid locket name: %v", err)
	}
	locket, err := s.Store.GetLocket(ctx, &store.FindLocket{
		ID: &id,
	})
	if err != nil {
		return nil, err
	}
	if locket == nil {
		return nil, status.Errorf(codes.NotFound, "locket not found")
	}

	user, _ := getCurrentUser(ctx, s.Store)
	if locket.CreatorID != user.ID {
		return nil, status.Errorf(codes.PermissionDenied, "permission denied")
	}

	if locketMessage, err := s.convertLocketFromStore(ctx, locket); err == nil {
		// Try to dispatch webhook when locket is deleted.
		if err := s.DispatchLocketDeletedWebhook(ctx, locketMessage); err != nil {
			slog.Warn("Failed to dispatch locket deleted webhook", err)
		}
	}

	if err = s.Store.DeleteLocket(ctx, &store.DeleteLocket{ID: id}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete locket")
	}

	return &apiv2pb.DeleteLocketResponse{}, nil
}

func (s *APIV2Service) CreateLocketComment(ctx context.Context, request *apiv2pb.CreateLocketCommentRequest) (*apiv2pb.CreateLocketCommentResponse, error) {
	id, err := ExtractLocketIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid locket name: %v", err)
	}
	relatedLocket, err := s.Store.GetLocket(ctx, &store.FindLocket{ID: &id})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get locket")
	}

	// Create the comment locket first.
	createLocketResponse, err := s.CreateLocket(ctx, request.Comment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create locket")
	}

	// Build the relation between the comment locket and the original locket.
	locket := createLocketResponse.Locket
	locketID, err := ExtractLocketIDFromName(locket.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid locket name: %v", err)
	}
	_, err = s.Store.UpsertLocketRelation(ctx, &store.LocketRelation{
		LocketID:        locketID,
		RelatedLocketID: relatedLocket.ID,
		Type:          store.LocketRelationComment,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create locket relation")
	}
	creatorID, err := ExtractUserIDFromName(locket.Creator)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid locket creator")
	}
	if locket.Visibility != apiv2pb.Visibility_PRIVATE && creatorID != relatedLocket.CreatorID {
		activity, err := s.Store.CreateActivity(ctx, &store.Activity{
			CreatorID: creatorID,
			Type:      store.ActivityTypeLocketComment,
			Level:     store.ActivityLevelInfo,
			Payload: &storepb.ActivityPayload{
				LocketComment: &storepb.ActivityLocketCommentPayload{
					LocketId:        locketID,
					RelatedLocketId: relatedLocket.ID,
				},
			},
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create activity")
		}
		if _, err := s.Store.CreateInbox(ctx, &store.Inbox{
			SenderID:   creatorID,
			ReceiverID: relatedLocket.CreatorID,
			Status:     store.UNREAD,
			Message: &storepb.InboxMessage{
				Type:       storepb.InboxMessage_TYPE_LOCKET_COMMENT,
				ActivityId: &activity.ID,
			},
		}); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create inbox")
		}
	}

	response := &apiv2pb.CreateLocketCommentResponse{
		Locket: locket,
	}
	return response, nil
}

func (s *APIV2Service) ListLocketComments(ctx context.Context, request *apiv2pb.ListLocketCommentsRequest) (*apiv2pb.ListLocketCommentsResponse, error) {
	id, err := ExtractLocketIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid locket name: %v", err)
	}
	locketRelationComment := store.LocketRelationComment
	locketRelations, err := s.Store.ListLocketRelations(ctx, &store.FindLocketRelation{
		RelatedLocketID: &id,
		Type:          &locketRelationComment,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list locket relations")
	}

	var lockets []*apiv2pb.Locket
	for _, locketRelation := range locketRelations {
		locket, err := s.Store.GetLocket(ctx, &store.FindLocket{
			ID: &locketRelation.LocketID,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get locket")
		}
		if locket != nil {
			locketMessage, err := s.convertLocketFromStore(ctx, locket)
			if err != nil {
				return nil, errors.Wrap(err, "failed to convert locket")
			}
			lockets = append(lockets, locketMessage)
		}
	}

	response := &apiv2pb.ListLocketCommentsResponse{
		Lockets: lockets,
	}
	return response, nil
}

func (s *APIV2Service) GetUserLocketsStats(ctx context.Context, request *apiv2pb.GetUserLocketsStatsRequest) (*apiv2pb.GetUserLocketsStatsResponse, error) {
	userID, err := ExtractUserIDFromName(request.Name)
	if err != nil {
		return nil, errors.Wrap(err, "invalid user name")
	}
	user, err := s.Store.GetUser(ctx, &store.FindUser{
		ID: &userID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user")
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	normalRowStatus := store.Normal
	locketFind := &store.FindLocket{
		CreatorID:       &user.ID,
		RowStatus:       &normalRowStatus,
		ExcludeComments: true,
		ExcludeContent:  true,
	}
	if err := s.buildLocketFindWithFilter(ctx, locketFind, request.Filter); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to build find lockets with filter")
	}

	lockets, err := s.Store.ListLockets(ctx, locketFind)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list lockets: %v", err)
	}

	location, err := time.LoadLocation(request.Timezone)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid timezone location")
	}

	displayWithUpdatedTs, err := s.getLocketDisplayWithUpdatedTsSettingValue(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get locket display with updated ts setting value")
	}
	stats := make(map[string]int32)
	for _, locket := range lockets {
		displayTs := locket.CreatedTs
		if displayWithUpdatedTs {
			displayTs = locket.UpdatedTs
		}
		stats[time.Unix(displayTs, 0).In(location).Format("2006-01-02")]++
	}

	response := &apiv2pb.GetUserLocketsStatsResponse{
		Stats: stats,
	}
	return response, nil
}

func (s *APIV2Service) ExportLockets(ctx context.Context, request *apiv2pb.ExportLocketsRequest) (*apiv2pb.ExportLocketsResponse, error) {
	normalRowStatus := store.Normal
	locketFind := &store.FindLocket{
		RowStatus: &normalRowStatus,
		// Exclude comments by default.
		ExcludeComments: true,
	}
	if err := s.buildLocketFindWithFilter(ctx, locketFind, request.Filter); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to build find lockets with filter: %v", err)
	}

	lockets, err := s.Store.ListLockets(ctx, locketFind)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list lockets: %v", err)
	}

	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	for _, locket := range lockets {
		locketMessage, err := s.convertLocketFromStore(ctx, locket)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert locket")
		}
		file, err := writer.Create(time.Unix(locket.CreatedTs, 0).Format(time.RFC3339) + ".md")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to create locket file")
		}
		_, err = file.Write([]byte(locketMessage.Content))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to write to locket file")
		}
	}
	if err := writer.Close(); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to close zip file writer")
	}

	return &apiv2pb.ExportLocketsResponse{
		Content: buf.Bytes(),
	}, nil
}

func (s *APIV2Service) convertLocketFromStore(ctx context.Context, locket *store.Locket) (*apiv2pb.Locket, error) {
	displayTs := locket.CreatedTs
	if displayWithUpdatedTs, err := s.getLocketDisplayWithUpdatedTsSettingValue(ctx); err == nil && displayWithUpdatedTs {
		displayTs = locket.UpdatedTs
	}

	creator, err := s.Store.GetUser(ctx, &store.FindUser{ID: &locket.CreatorID})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get creator")
	}

	name := fmt.Sprintf("%s%d", LocketNamePrefix, locket.ID)
	listLocketRelationsResponse, err := s.ListLocketRelations(ctx, &apiv2pb.ListLocketRelationsRequest{Name: name})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list locket relations")
	}

	listLocketResourcesResponse, err := s.ListLocketResources(ctx, &apiv2pb.ListLocketResourcesRequest{Name: name})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list locket resources")
	}

	listLocketReactionsResponse, err := s.ListLocketReactions(ctx, &apiv2pb.ListLocketReactionsRequest{Name: name})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list locket reactions")
	}

	return &apiv2pb.Locket{
		Name:        name,
		Uid:         locket.UID,
		RowStatus:   convertRowStatusFromStore(locket.RowStatus),
		Creator:     fmt.Sprintf("%s%d", UserNamePrefix, creator.ID),
		CreateTime:  timestamppb.New(time.Unix(locket.CreatedTs, 0)),
		UpdateTime:  timestamppb.New(time.Unix(locket.UpdatedTs, 0)),
		DisplayTime: timestamppb.New(time.Unix(displayTs, 0)),
		Content:     locket.Content,
		Visibility:  convertVisibilityFromStore(locket.Visibility),
		Pinned:      locket.Pinned,
		ParentId:    locket.ParentID,
		Relations:   listLocketRelationsResponse.Relations,
		Resources:   listLocketResourcesResponse.Resources,
		Reactions:   listLocketReactionsResponse.Reactions,
	}, nil
}

func (s *APIV2Service) getLocketDisplayWithUpdatedTsSettingValue(ctx context.Context) (bool, error) {
	locketDisplayWithUpdatedTsSetting, err := s.Store.GetWorkspaceSetting(ctx, &store.FindWorkspaceSetting{
		Name: apiv1.SystemSettingLocketDisplayWithUpdatedTsName.String(),
	})
	if err != nil {
		return false, errors.Wrap(err, "failed to find system setting")
	}
	if locketDisplayWithUpdatedTsSetting == nil {
		return false, nil
	}

	locketDisplayWithUpdatedTs := false
	if err := json.Unmarshal([]byte(locketDisplayWithUpdatedTsSetting.Value), &locketDisplayWithUpdatedTs); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal system setting value")
	}
	return locketDisplayWithUpdatedTs, nil
}

func (s *APIV2Service) getDisablePublicLocketsSystemSettingValue(ctx context.Context) (bool, error) {
	disablePublicLocketsSystemSetting, err := s.Store.GetWorkspaceSetting(ctx, &store.FindWorkspaceSetting{
		Name: apiv1.SystemSettingDisablePublicLocketsName.String(),
	})
	if err != nil {
		return false, errors.Wrap(err, "failed to find system setting")
	}
	if disablePublicLocketsSystemSetting == nil {
		return false, nil
	}

	disablePublicLockets := false
	if err := json.Unmarshal([]byte(disablePublicLocketsSystemSetting.Value), &disablePublicLockets); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal system setting value")
	}
	return disablePublicLockets, nil
}

func convertVisibilityFromStore(visibility store.Visibility) apiv2pb.Visibility {
	switch visibility {
	case store.Private:
		return apiv2pb.Visibility_PRIVATE
	case store.Protected:
		return apiv2pb.Visibility_PROTECTED
	case store.Public:
		return apiv2pb.Visibility_PUBLIC
	default:
		return apiv2pb.Visibility_VISIBILITY_UNSPECIFIED
	}
}

func convertVisibilityToStore(visibility apiv2pb.Visibility) store.Visibility {
	switch visibility {
	case apiv2pb.Visibility_PRIVATE:
		return store.Private
	case apiv2pb.Visibility_PROTECTED:
		return store.Protected
	case apiv2pb.Visibility_PUBLIC:
		return store.Public
	default:
		return store.Private
	}
}

func (s *APIV2Service) buildLocketFindWithFilter(ctx context.Context, find *store.FindLocket, filter string) error {
	user, _ := getCurrentUser(ctx, s.Store)
	if find == nil {
		find = &store.FindLocket{}
	}
	if filter != "" {
		filter, err := parseSearchLocketsFilter(filter)
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "invalid filter: %v", err)
		}
		if len(filter.ContentSearch) > 0 {
			find.ContentSearch = filter.ContentSearch
		}
		if len(filter.Visibilities) > 0 {
			find.VisibilityList = filter.Visibilities
		}
		if filter.OrderByPinned {
			find.OrderByPinned = filter.OrderByPinned
		}
		if filter.DisplayTimeAfter != nil {
			displayWithUpdatedTs, err := s.getLocketDisplayWithUpdatedTsSettingValue(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "failed to get locket display with updated ts setting value")
			}
			if displayWithUpdatedTs {
				find.UpdatedTsAfter = filter.DisplayTimeAfter
			} else {
				find.CreatedTsAfter = filter.DisplayTimeAfter
			}
		}
		if filter.DisplayTimeBefore != nil {
			displayWithUpdatedTs, err := s.getLocketDisplayWithUpdatedTsSettingValue(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "failed to get locket display with updated ts setting value")
			}
			if displayWithUpdatedTs {
				find.UpdatedTsBefore = filter.DisplayTimeBefore
			} else {
				find.CreatedTsBefore = filter.DisplayTimeBefore
			}
		}
		if filter.Creator != nil {
			userID, err := ExtractUserIDFromName(*filter.Creator)
			if err != nil {
				return errors.Wrap(err, "invalid user name")
			}
			user, err := s.Store.GetUser(ctx, &store.FindUser{
				ID: &userID,
			})
			if err != nil {
				return status.Errorf(codes.Internal, "failed to get user")
			}
			if user == nil {
				return status.Errorf(codes.NotFound, "user not found")
			}
			find.CreatorID = &user.ID
		}
		if filter.UID != nil {
			find.UID = filter.UID
		}
		if filter.RowStatus != nil {
			find.RowStatus = filter.RowStatus
		}
		if filter.Random {
			find.Random = filter.Random
		}
		if filter.Limit != nil {
			find.Limit = filter.Limit
		}
	}

	// If the user is not authenticated, only public lockets are visible.
	if user == nil {
		if filter == "" {
			// If no filter is provided, return an error.
			return status.Errorf(codes.InvalidArgument, "filter is required")
		}

		find.VisibilityList = []store.Visibility{store.Public}
	} else if find.CreatorID != nil && *find.CreatorID != user.ID {
		find.VisibilityList = []store.Visibility{store.Public, store.Protected}
	}

	displayWithUpdatedTs, err := s.getLocketDisplayWithUpdatedTsSettingValue(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to get locket display with updated ts setting value")
	}
	if displayWithUpdatedTs {
		find.OrderByUpdatedTs = true
	}
	return nil
}

// SearchLocketsFilterCELAttributes are the CEL attributes.
var SearchLocketsFilterCELAttributes = []cel.EnvOption{
	cel.Variable("content_search", cel.ListType(cel.StringType)),
	cel.Variable("visibilities", cel.ListType(cel.StringType)),
	cel.Variable("order_by_pinned", cel.BoolType),
	cel.Variable("display_time_before", cel.IntType),
	cel.Variable("display_time_after", cel.IntType),
	cel.Variable("creator", cel.StringType),
	cel.Variable("uid", cel.StringType),
	cel.Variable("row_status", cel.StringType),
	cel.Variable("random", cel.BoolType),
	cel.Variable("limit", cel.IntType),
}

type SearchLocketsFilter struct {
	ContentSearch     []string
	Visibilities      []store.Visibility
	OrderByPinned     bool
	DisplayTimeBefore *int64
	DisplayTimeAfter  *int64
	Creator           *string
	UID               *string
	RowStatus         *store.RowStatus
	Random            bool
	Limit             *int
}

func parseSearchLocketsFilter(expression string) (*SearchLocketsFilter, error) {
	e, err := cel.NewEnv(SearchLocketsFilterCELAttributes...)
	if err != nil {
		return nil, err
	}
	ast, issues := e.Compile(expression)
	if issues != nil {
		return nil, errors.Errorf("found issue %v", issues)
	}
	filter := &SearchLocketsFilter{}
	expr, err := cel.AstToParsedExpr(ast)
	if err != nil {
		return nil, err
	}
	callExpr := expr.GetExpr().GetCallExpr()
	findSearchLocketsField(callExpr, filter)
	return filter, nil
}

func findSearchLocketsField(callExpr *expr.Expr_Call, filter *SearchLocketsFilter) {
	if len(callExpr.Args) == 2 {
		idExpr := callExpr.Args[0].GetIdentExpr()
		if idExpr != nil {
			if idExpr.Name == "content_search" {
				contentSearch := []string{}
				for _, expr := range callExpr.Args[1].GetListExpr().GetElements() {
					value := expr.GetConstExpr().GetStringValue()
					contentSearch = append(contentSearch, value)
				}
				filter.ContentSearch = contentSearch
			} else if idExpr.Name == "visibilities" {
				visibilities := []store.Visibility{}
				for _, expr := range callExpr.Args[1].GetListExpr().GetElements() {
					value := expr.GetConstExpr().GetStringValue()
					visibilities = append(visibilities, store.Visibility(value))
				}
				filter.Visibilities = visibilities
			} else if idExpr.Name == "order_by_pinned" {
				value := callExpr.Args[1].GetConstExpr().GetBoolValue()
				filter.OrderByPinned = value
			} else if idExpr.Name == "display_time_before" {
				displayTimeBefore := callExpr.Args[1].GetConstExpr().GetInt64Value()
				filter.DisplayTimeBefore = &displayTimeBefore
			} else if idExpr.Name == "display_time_after" {
				displayTimeAfter := callExpr.Args[1].GetConstExpr().GetInt64Value()
				filter.DisplayTimeAfter = &displayTimeAfter
			} else if idExpr.Name == "creator" {
				creator := callExpr.Args[1].GetConstExpr().GetStringValue()
				filter.Creator = &creator
			} else if idExpr.Name == "uid" {
				uid := callExpr.Args[1].GetConstExpr().GetStringValue()
				filter.UID = &uid
			} else if idExpr.Name == "row_status" {
				rowStatus := store.RowStatus(callExpr.Args[1].GetConstExpr().GetStringValue())
				filter.RowStatus = &rowStatus
			} else if idExpr.Name == "random" {
				value := callExpr.Args[1].GetConstExpr().GetBoolValue()
				filter.Random = value
			} else if idExpr.Name == "limit" {
				limit := int(callExpr.Args[1].GetConstExpr().GetInt64Value())
				filter.Limit = &limit
			}
			return
		}
	}
	for _, arg := range callExpr.Args {
		callExpr := arg.GetCallExpr()
		if callExpr != nil {
			findSearchLocketsField(callExpr, filter)
		}
	}
}

// DispatchLocketCreatedWebhook dispatches webhook when locket is created.
func (s *APIV2Service) DispatchLocketCreatedWebhook(ctx context.Context, locket *apiv2pb.Locket) error {
	return s.dispatchLocketRelatedWebhook(ctx, locket, "lockets.locket.created")
}

// DispatchLocketUpdatedWebhook dispatches webhook when locket is updated.
func (s *APIV2Service) DispatchLocketUpdatedWebhook(ctx context.Context, locket *apiv2pb.Locket) error {
	return s.dispatchLocketRelatedWebhook(ctx, locket, "lockets.locket.updated")
}

// DispatchLocketDeletedWebhook dispatches webhook when locket is deleted.
func (s *APIV2Service) DispatchLocketDeletedWebhook(ctx context.Context, locket *apiv2pb.Locket) error {
	return s.dispatchLocketRelatedWebhook(ctx, locket, "lockets.locket.deleted")
}

func (s *APIV2Service) dispatchLocketRelatedWebhook(ctx context.Context, locket *apiv2pb.Locket, activityType string) error {
	creatorID, err := ExtractUserIDFromName(locket.Creator)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid locket creator")
	}
	webhooks, err := s.Store.ListWebhooks(ctx, &store.FindWebhook{
		CreatorID: &creatorID,
	})
	if err != nil {
		return err
	}
	for _, hook := range webhooks {
		payload, err := convertLocketToWebhookPayload(locket)
		if err != nil {
			return errors.Wrap(err, "failed to convert locket to webhook payload")
		}
		payload.ActivityType = activityType
		payload.URL = hook.Url
		if err := webhook.Post(*payload); err != nil {
			return errors.Wrap(err, "failed to post webhook")
		}
	}
	return nil
}

func convertLocketToWebhookPayload(locket *apiv2pb.Locket) (*webhook.WebhookPayload, error) {
	creatorID, err := ExtractUserIDFromName(locket.Creator)
	if err != nil {
		return nil, errors.Wrap(err, "invalid locket creator")
	}
	id, err := ExtractLocketIDFromName(locket.Name)
	if err != nil {
		return nil, errors.Wrap(err, "invalid locket name")
	}
	return &webhook.WebhookPayload{
		CreatorID: creatorID,
		CreatedTs: time.Now().Unix(),
		Locket: &webhook.Locket{
			ID:         id,
			CreatorID:  creatorID,
			CreatedTs:  locket.CreateTime.Seconds,
			UpdatedTs:  locket.UpdateTime.Seconds,
			Content:    locket.Content,
			Visibility: locket.Visibility.String(),
			Pinned:     locket.Pinned,
			ResourceList: func() []*webhook.Resource {
				resources := []*webhook.Resource{}
				for _, resource := range locket.Resources {
					resources = append(resources, &webhook.Resource{
						UID:          resource.Uid,
						Filename:     resource.Filename,
						ExternalLink: resource.ExternalLink,
						Type:         resource.Type,
						Size:         resource.Size,
					})
				}
				return resources
			}(),
		},
	}, nil
}
