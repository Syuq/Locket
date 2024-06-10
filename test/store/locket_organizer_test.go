package teststore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Syuq/Locket/store"
)

func TestLocketOrganizerStore(t *testing.T) {
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

	locketOrganizer, err := ts.UpsertLocketOrganizer(ctx, &store.LocketOrganizer{
		LocketID: locket.ID,
		UserID: user.ID,
		Pinned: true,
	})
	require.NoError(t, err)
	require.NotNil(t, locketOrganizer)
	require.Equal(t, locket.ID, locketOrganizer.LocketID)
	require.Equal(t, user.ID, locketOrganizer.UserID)
	require.Equal(t, true, locketOrganizer.Pinned)

	locketOrganizerTemp, err := ts.GetLocketOrganizer(ctx, &store.FindLocketOrganizer{
		LocketID: locket.ID,
	})
	require.NoError(t, err)
	require.Equal(t, locketOrganizer, locketOrganizerTemp)
	locketOrganizerTemp, err = ts.UpsertLocketOrganizer(ctx, &store.LocketOrganizer{
		LocketID: locket.ID,
		UserID: user.ID,
		Pinned: false,
	})
	require.NoError(t, err)
	require.NotNil(t, locketOrganizerTemp)
	require.Equal(t, locket.ID, locketOrganizerTemp.LocketID)
	require.Equal(t, user.ID, locketOrganizerTemp.UserID)
	require.Equal(t, false, locketOrganizerTemp.Pinned)
	err = ts.DeleteLocketOrganizer(ctx, &store.DeleteLocketOrganizer{
		LocketID: &locket.ID,
		UserID: &user.ID,
	})
	require.NoError(t, err)
	locketOrganizers, err := ts.ListLocketOrganizer(ctx, &store.FindLocketOrganizer{
		UserID: user.ID,
	})
	require.NoError(t, err)
	require.Equal(t, 0, len(locketOrganizers))
	ts.Close()
}
