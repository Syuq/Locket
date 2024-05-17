package store

import (
	"context"
	"errors"
)

type LocketOrganizer struct {
	LocketID int32
	UserID int32
	Pinned bool
}

type FindLocketOrganizer struct {
	LocketID int32
	UserID int32
}

type DeleteLocketOrganizer struct {
	LocketID *int32
	UserID *int32
}

func (s *Store) UpsertLocketOrganizer(ctx context.Context, upsert *LocketOrganizer) (*LocketOrganizer, error) {
	return s.driver.UpsertLocketOrganizer(ctx, upsert)
}

func (s *Store) GetLocketOrganizer(ctx context.Context, find *FindLocketOrganizer) (*LocketOrganizer, error) {
	list, err := s.ListLocketOrganizer(ctx, find)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, errors.New("not found")
	}

	return list[0], nil
}

func (s *Store) ListLocketOrganizer(ctx context.Context, find *FindLocketOrganizer) ([]*LocketOrganizer, error) {
	return s.driver.ListLocketOrganizer(ctx, find)
}

func (s *Store) DeleteLocketOrganizer(ctx context.Context, delete *DeleteLocketOrganizer) error {
	return s.driver.DeleteLocketOrganizer(ctx, delete)
}
