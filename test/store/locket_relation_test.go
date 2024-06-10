package teststore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Syuq/Locket/store"
)

func TestLocketRelationStore(t *testing.T) {
	ctx := context.Background()
	ts := NewTestingStore(ctx, t)
	user, err := createTestingHostUser(ctx, ts)
	require.NoError(t, err)
	locketCreate := &store.Locket{
		UID:        "main-locket",
		CreatorID:  user.ID,
		Content:    "main locket content",
		Visibility: store.Public,
	}
	locket, err := ts.CreateLocket(ctx, locketCreate)
	require.NoError(t, err)
	require.Equal(t, locketCreate.Content, locket.Content)
	relatedLocketCreate := &store.Locket{
		UID:        "related-locket",
		CreatorID:  user.ID,
		Content:    "related locket content",
		Visibility: store.Public,
	}
	relatedLocket, err := ts.CreateLocket(ctx, relatedLocketCreate)
	require.NoError(t, err)
	require.Equal(t, relatedLocketCreate.Content, relatedLocket.Content)
	commentLocketCreate := &store.Locket{
		UID:        "comment-locket",
		CreatorID:  user.ID,
		Content:    "comment locket content",
		Visibility: store.Public,
	}
	commentLocket, err := ts.CreateLocket(ctx, commentLocketCreate)
	require.NoError(t, err)
	require.Equal(t, commentLocketCreate.Content, commentLocket.Content)

	// Reference relation.
	referenceRelation := &store.LocketRelation{
		LocketID:        locket.ID,
		RelatedLocketID: relatedLocket.ID,
		Type:          store.LocketRelationReference,
	}
	_, err = ts.UpsertLocketRelation(ctx, referenceRelation)
	require.NoError(t, err)
	// Comment relation.
	commentRelation := &store.LocketRelation{
		LocketID:        locket.ID,
		RelatedLocketID: commentLocket.ID,
		Type:          store.LocketRelationComment,
	}
	_, err = ts.UpsertLocketRelation(ctx, commentRelation)
	require.NoError(t, err)
	ts.Close()
}
