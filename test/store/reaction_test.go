package teststore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	storepb "github.com/Syuq/Locket/proto/gen/store"
	"github.com/Syuq/Locket/store"
)

func TestReactionStore(t *testing.T) {
	ctx := context.Background()
	ts := NewTestingStore(ctx, t)

	user, err := createTestingHostUser(ctx, ts)
	require.NoError(t, err)

	contentID := "test_content_id"
	reaction, err := ts.UpsertReaction(ctx, &storepb.Reaction{
		CreatorId:    user.ID,
		ContentId:    contentID,
		ReactionType: storepb.Reaction_HEART,
	})
	require.NoError(t, err)
	require.NotNil(t, reaction)
	require.NotEmpty(t, reaction.Id)

	reactions, err := ts.ListReactions(ctx, &store.FindReaction{
		ContentID: &contentID,
	})
	require.NoError(t, err)
	require.Len(t, reactions, 1)
	require.Equal(t, reaction, reactions[0])

	err = ts.DeleteReaction(ctx, &store.DeleteReaction{
		ID: reaction.Id,
	})
	require.NoError(t, err)

	reactions, err = ts.ListReactions(ctx, &store.FindReaction{
		ContentID: &contentID,
	})
	require.NoError(t, err)
	require.Len(t, reactions, 0)

	ts.Close()
}
