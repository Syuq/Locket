package teststore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Syuq/Locket/store"
)

func TestLocketStore(t *testing.T) {
	ctx := context.Background()
	ts := NewTestingStore(ctx, t)
	user, err := createTestingHostUser(ctx, ts)
	require.NoError(t, err)
	locketCreate := &store.Locket{
		UID:        "test-resource-name",
		CreatorID:  user.ID,
		Content:    "test_content",
		Visibility: store.Public,
	}
	locket, err := ts.CreateLocket(ctx, locketCreate)
	require.NoError(t, err)
	require.Equal(t, locketCreate.Content, locket.Content)
	locketPatchContent := "test_content_2"
	locketPatch := &store.UpdateLocket{
		ID:      locket.ID,
		Content: &locketPatchContent,
	}
	err = ts.UpdateLocket(ctx, locketPatch)
	require.NoError(t, err)
	locket, err = ts.GetLocket(ctx, &store.FindLocket{
		ID: &locket.ID,
	})
	require.NoError(t, err)
	require.NotNil(t, locket)
	locketList, err := ts.ListLockets(ctx, &store.FindLocket{
		CreatorID: &user.ID,
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(locketList))
	require.Equal(t, locket, locketList[0])
	err = ts.DeleteLocket(ctx, &store.DeleteLocket{
		ID: locket.ID,
	})
	require.NoError(t, err)
	locketList, err = ts.ListLockets(ctx, &store.FindLocket{
		CreatorID: &user.ID,
	})
	require.NoError(t, err)
	require.Equal(t, 0, len(locketList))

	locketList, err = ts.ListLockets(ctx, &store.FindLocket{
		CreatorID: &user.ID,
		VisibilityList: []store.Visibility{
			store.Public,
		},
	})
	require.NoError(t, err)
	require.Equal(t, 0, len(locketList))
	ts.Close()
}

func TestDeleteLocketStore(t *testing.T) {
	ctx := context.Background()
	ts := NewTestingStore(ctx, t)
	user, err := createTestingHostUser(ctx, ts)
	require.NoError(t, err)
	locketCreate := &store.Locket{
		UID:        "test-resource-name",
		CreatorID:  user.ID,
		Content:    "test_content",
		Visibility: store.Public,
	}
	locket, err := ts.CreateLocket(ctx, locketCreate)
	require.NoError(t, err)
	require.Equal(t, locketCreate.Content, locket.Content)
	err = ts.DeleteLocket(ctx, &store.DeleteLocket{
		ID: locket.ID,
	})
	require.NoError(t, err)
	ts.Close()
}
