package store

import (
	"context"
	"errors"

	"github.com/syuq/locket/internal/util"
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
	// ID is the system generated unique identifier for the locket.
	ID int32
	// UID is the user defined unique identifier for the locket.
	UID string

	// Standard fields
	RowStatus RowStatus
	CreatorID int32
	CreatedTs int64
	UpdatedTs int64

	// Domain specific fields
	Content    string
	Visibility Visibility

	// Composed fields
	Pinned   bool
	ParentID *int32
}

type FindLocket struct {
	ID  *int32
	UID *string

	// Standard fields
	RowStatus       *RowStatus
	CreatorID       *int32
	CreatedTsAfter  *int64
	CreatedTsBefore *int64
	UpdatedTsAfter  *int64
	UpdatedTsBefore *int64

	// Domain specific fields
	ContentSearch   []string
	VisibilityList  []Visibility
	ExcludeContent  bool
	ExcludeComments bool
	Random          bool

	// Pagination
	Limit            *int
	Offset           *int
	OrderByUpdatedTs bool
	OrderByPinned    bool
}

type UpdateLocket struct {
	ID         int32
	UID        *string
	CreatedTs  *int64
	UpdatedTs  *int64
	RowStatus  *RowStatus
	Content    *string
	Visibility *Visibility
}

type DeleteLocket struct {
	ID int32
}

func (s *Store) CreateLocket(ctx context.Context, create *Locket) (*Locket, error) {
	if !util.UIDMatcher.MatchString(create.UID) {
		return nil, errors.New("invalid uid")
	}
	return s.driver.CreateLocket(ctx, create)
}

func (s *Store) ListLockets(ctx context.Context, find *FindLocket) ([]*Locket, error) {
	return s.driver.ListLockets(ctx, find)
}

func (s *Store) GetLocket(ctx context.Context, find *FindLocket) (*Locket, error) {
	list, err := s.ListLockets(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}

	locket := list[0]
	return locket, nil
}

func (s *Store) UpdateLocket(ctx context.Context, update *UpdateLocket) error {
	if update.UID != nil && !util.UIDMatcher.MatchString(*update.UID) {
		return errors.New("invalid uid")
	}
	return s.driver.UpdateLocket(ctx, update)
}

func (s *Store) DeleteLocket(ctx context.Context, delete *DeleteLocket) error {
	return s.driver.DeleteLocket(ctx, delete)
}
