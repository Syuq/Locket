package v2

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apiv2pb "github.com/syuq/locket/proto/gen/api/v2"
	"github.com/syuq/locket/store"
)

func (s *APIV2Service) SetLocketRelations(ctx context.Context, request *apiv2pb.SetLocketRelationsRequest) (*apiv2pb.SetLocketRelationsResponse, error) {
	id, err := ExtractLocketIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid locket name: %v", err)
	}
	referenceType := store.LocketRelationReference
	// Delete all reference relations first.
	if err := s.Store.DeleteLocketRelation(ctx, &store.DeleteLocketRelation{
		LocketID: &id,
		Type:   &referenceType,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete locket relation")
	}

	for _, relation := range request.Relations {
		// Ignore reflexive relations.
		if request.Name == relation.RelatedLocket {
			continue
		}
		// Ignore comment relations as there's no need to update a comment's relation.
		// Inserting/Deleting a comment is handled elsewhere.
		if relation.Type == apiv2pb.LocketRelation_COMMENT {
			continue
		}
		relatedLocketID, err := ExtractLocketIDFromName(relation.RelatedLocket)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid related locket name: %v", err)
		}
		if _, err := s.Store.UpsertLocketRelation(ctx, &store.LocketRelation{
			LocketID:        id,
			RelatedLocketID: relatedLocketID,
			Type:          convertLocketRelationTypeToStore(relation.Type),
		}); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to upsert locket relation")
		}
	}

	return &apiv2pb.SetLocketRelationsResponse{}, nil
}

func (s *APIV2Service) ListLocketRelations(ctx context.Context, request *apiv2pb.ListLocketRelationsRequest) (*apiv2pb.ListLocketRelationsResponse, error) {
	id, err := ExtractLocketIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid locket name: %v", err)
	}
	relationList := []*apiv2pb.LocketRelation{}
	tempList, err := s.Store.ListLocketRelations(ctx, &store.FindLocketRelation{
		LocketID: &id,
	})
	if err != nil {
		return nil, err
	}
	for _, relation := range tempList {
		relationList = append(relationList, convertLocketRelationFromStore(relation))
	}
	tempList, err = s.Store.ListLocketRelations(ctx, &store.FindLocketRelation{
		RelatedLocketID: &id,
	})
	if err != nil {
		return nil, err
	}
	for _, relation := range tempList {
		relationList = append(relationList, convertLocketRelationFromStore(relation))
	}

	response := &apiv2pb.ListLocketRelationsResponse{
		Relations: relationList,
	}
	return response, nil
}

func convertLocketRelationFromStore(locketRelation *store.LocketRelation) *apiv2pb.LocketRelation {
	return &apiv2pb.LocketRelation{
		Locket:        fmt.Sprintf("%s%d", LocketNamePrefix, locketRelation.LocketID),
		RelatedLocket: fmt.Sprintf("%s%d", LocketNamePrefix, locketRelation.RelatedLocketID),
		Type:        convertLocketRelationTypeFromStore(locketRelation.Type),
	}
}

func convertLocketRelationTypeFromStore(relationType store.LocketRelationType) apiv2pb.LocketRelation_Type {
	switch relationType {
	case store.LocketRelationReference:
		return apiv2pb.LocketRelation_REFERENCE
	case store.LocketRelationComment:
		return apiv2pb.LocketRelation_COMMENT
	default:
		return apiv2pb.LocketRelation_TYPE_UNSPECIFIED
	}
}

func convertLocketRelationTypeToStore(relationType apiv2pb.LocketRelation_Type) store.LocketRelationType {
	switch relationType {
	case apiv2pb.LocketRelation_REFERENCE:
		return store.LocketRelationReference
	case apiv2pb.LocketRelation_COMMENT:
		return store.LocketRelationComment
	default:
		return store.LocketRelationReference
	}
}
