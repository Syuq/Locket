package v2

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apiv2pb "github.com/Syuq/Locket/proto/gen/api/v2"
	storepb "github.com/Syuq/Locket/proto/gen/store"
	"github.com/Syuq/Locket/store"
)

func (s *APIV2Service) ListLocketReactions(ctx context.Context, request *apiv2pb.ListLocketReactionsRequest) (*apiv2pb.ListLocketReactionsResponse, error) {
	reactions, err := s.Store.ListReactions(ctx, &store.FindReaction{
		ContentID: &request.Name,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list reactions")
	}

	response := &apiv2pb.ListLocketReactionsResponse{
		Reactions: []*apiv2pb.Reaction{},
	}
	for _, reaction := range reactions {
		reactionMessage, err := s.convertReactionFromStore(ctx, reaction)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to convert reaction")
		}
		response.Reactions = append(response.Reactions, reactionMessage)
	}
	return response, nil
}

func (s *APIV2Service) UpsertLocketReaction(ctx context.Context, request *apiv2pb.UpsertLocketReactionRequest) (*apiv2pb.UpsertLocketReactionResponse, error) {
	user, err := getCurrentUser(ctx, s.Store)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get current user")
	}
	reaction, err := s.Store.UpsertReaction(ctx, &storepb.Reaction{
		CreatorId:    user.ID,
		ContentId:    request.Reaction.ContentId,
		ReactionType: storepb.Reaction_Type(request.Reaction.ReactionType),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to upsert reaction")
	}

	reactionMessage, err := s.convertReactionFromStore(ctx, reaction)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert reaction")
	}
	return &apiv2pb.UpsertLocketReactionResponse{
		Reaction: reactionMessage,
	}, nil
}

func (s *APIV2Service) DeleteLocketReaction(ctx context.Context, request *apiv2pb.DeleteLocketReactionRequest) (*apiv2pb.DeleteLocketReactionResponse, error) {
	if err := s.Store.DeleteReaction(ctx, &store.DeleteReaction{
		ID: request.ReactionId,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete reaction")
	}

	return &apiv2pb.DeleteLocketReactionResponse{}, nil
}

func (s *APIV2Service) convertReactionFromStore(ctx context.Context, reaction *storepb.Reaction) (*apiv2pb.Reaction, error) {
	creator, err := s.Store.GetUser(ctx, &store.FindUser{
		ID: &reaction.CreatorId,
	})
	if err != nil {
		return nil, err
	}
	return &apiv2pb.Reaction{
		Id:           reaction.Id,
		Creator:      fmt.Sprintf("%s%d", UserNamePrefix, creator.ID),
		ContentId:    reaction.ContentId,
		ReactionType: apiv2pb.Reaction_Type(reaction.ReactionType),
	}, nil
}
