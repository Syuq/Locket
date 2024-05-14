package v2

import (
	"context"
	"slices"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apiv2pb "github.com/syuq/locket/proto/gen/api/v2"
	"github.com/syuq/locket/store"
)

func (s *APIV2Service) SetLocketResources(ctx context.Context, request *apiv2pb.SetLocketResourcesRequest) (*apiv2pb.SetLocketResourcesResponse, error) {
	locketID, err := ExtractLocketIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid locket name: %v", err)
	}
	resources, err := s.Store.ListResources(ctx, &store.FindResource{
		LocketID: &locketID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list resources")
	}

	// Delete resources that are not in the request.
	for _, resource := range resources {
		found := false
		for _, requestResource := range request.Resources {
			if resource.UID == requestResource.Uid {
				found = true
				break
			}
		}
		if !found {
			if err = s.Store.DeleteResource(ctx, &store.DeleteResource{
				ID:     int32(resource.ID),
				LocketID: &locketID,
			}); err != nil {
				return nil, status.Errorf(codes.Internal, "failed to delete resource")
			}
		}
	}

	slices.Reverse(request.Resources)
	// Update resources' locket_id in the request.
	for index, resource := range request.Resources {
		id, err := ExtractResourceIDFromName(resource.Name)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid resource name: %v", err)
		}
		updatedTs := time.Now().Unix() + int64(index)
		if _, err := s.Store.UpdateResource(ctx, &store.UpdateResource{
			ID:        id,
			LocketID:    &locketID,
			UpdatedTs: &updatedTs,
		}); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update resource: %v", err)
		}
	}

	return &apiv2pb.SetLocketResourcesResponse{}, nil
}

func (s *APIV2Service) ListLocketResources(ctx context.Context, request *apiv2pb.ListLocketResourcesRequest) (*apiv2pb.ListLocketResourcesResponse, error) {
	id, err := ExtractLocketIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid locket name: %v", err)
	}
	resources, err := s.Store.ListResources(ctx, &store.FindResource{
		LocketID: &id,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list resources")
	}

	response := &apiv2pb.ListLocketResourcesResponse{
		Resources: []*apiv2pb.Resource{},
	}
	for _, resource := range resources {
		response.Resources = append(response.Resources, s.convertResourceFromStore(ctx, resource))
	}
	return response, nil
}
