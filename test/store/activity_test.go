package teststore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	storepb "github.com/syuq/locket/proto/gen/store"
	"github.com/syuq/locket/store"
)

func TestActivityStore(t *testing.T) {
	ctx := context.Background()
	ts := NewTestingStore(ctx, t)
	user, err := createTestingHostUser(ctx, ts)
	require.NoError(t, err)
	create := &store.Activity{
		CreatorID: user.ID,
		Type:      store.ActivityTypeLocketComment,
		Level:     store.ActivityLevelInfo,
		Payload:   &storepb.ActivityPayload{},
	}
	activity, err := ts.CreateActivity(ctx, create)
	require.NoError(t, err)
	require.NotNil(t, activity)
	activities, err := ts.ListActivities(ctx, &store.FindActivity{
		ID: &activity.ID,
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(activities))
	require.Equal(t, activity, activities[0])
	ts.Close()
}
