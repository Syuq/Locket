package store

import (
	"context"
)

type LocketRelationType string

const (
	// LocketRelationReference is the type for a reference locket relation.
	LocketRelationReference LocketRelationType = "REFERENCE"
	// LocketRelationComment is the type for a comment locket relation.
	LocketRelationComment LocketRelationType = "COMMENT"
)

type LocketRelation struct {
	LocketID        int32
	RelatedLocketID int32
	Type          LocketRelationType
}

type FindLocketRelation struct {
	LocketID        *int32
	RelatedLocketID *int32
	Type          *LocketRelationType
}

type DeleteLocketRelation struct {
	LocketID        *int32
	RelatedLocketID *int32
	Type          *LocketRelationType
}

func (s *Store) UpsertLocketRelation(ctx context.Context, create *LocketRelation) (*LocketRelation, error) {
	return s.driver.UpsertLocketRelation(ctx, create)
}

func (s *Store) ListLocketRelations(ctx context.Context, find *FindLocketRelation) ([]*LocketRelation, error) {
	return s.driver.ListLocketRelations(ctx, find)
}

func (s *Store) DeleteLocketRelation(ctx context.Context, delete *DeleteLocketRelation) error {
	return s.driver.DeleteLocketRelation(ctx, delete)
}
