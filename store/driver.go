package store

import (
	"context"
	"database/sql"

	storepb "github.com/Syuq/Locket/proto/gen/store"
)

// Driver is an interface for store driver.
// It contains all methods that store database driver should implement.
type Driver interface {
	GetDB() *sql.DB
	Close() error

	Migrate(ctx context.Context) error
	Vacuum(ctx context.Context) error

	// current file is driver
	GetCurrentDBSize(ctx context.Context) (int64, error)

	// MigrationHistory model related methods.
	FindMigrationHistoryList(ctx context.Context, find *FindMigrationHistory) ([]*MigrationHistory, error)
	UpsertMigrationHistory(ctx context.Context, upsert *UpsertMigrationHistory) (*MigrationHistory, error)

	// Activity model related methods.
	CreateActivity(ctx context.Context, create *Activity) (*Activity, error)
	ListActivities(ctx context.Context, find *FindActivity) ([]*Activity, error)

	// Resource model related methods.
	CreateResource(ctx context.Context, create *Resource) (*Resource, error)
	ListResources(ctx context.Context, find *FindResource) ([]*Resource, error)
	UpdateResource(ctx context.Context, update *UpdateResource) (*Resource, error)
	DeleteResource(ctx context.Context, delete *DeleteResource) error

	// Locket model related methods.
	CreateLocket(ctx context.Context, create *Locket) (*Locket, error)
	ListLockets(ctx context.Context, find *FindLocket) ([]*Locket, error)
	UpdateLocket(ctx context.Context, update *UpdateLocket) error
	DeleteLocket(ctx context.Context, delete *DeleteLocket) error

	// LocketRelation model related methods.
	UpsertLocketRelation(ctx context.Context, create *LocketRelation) (*LocketRelation, error)
	ListLocketRelations(ctx context.Context, find *FindLocketRelation) ([]*LocketRelation, error)
	DeleteLocketRelation(ctx context.Context, delete *DeleteLocketRelation) error

	// LocketOrganizer model related methods.
	UpsertLocketOrganizer(ctx context.Context, upsert *LocketOrganizer) (*LocketOrganizer, error)
	ListLocketOrganizer(ctx context.Context, find *FindLocketOrganizer) ([]*LocketOrganizer, error)
	DeleteLocketOrganizer(ctx context.Context, delete *DeleteLocketOrganizer) error

	// WorkspaceSetting model related methods.
	UpsertWorkspaceSetting(ctx context.Context, upsert *WorkspaceSetting) (*WorkspaceSetting, error)
	ListWorkspaceSettings(ctx context.Context, find *FindWorkspaceSetting) ([]*WorkspaceSetting, error)
	DeleteWorkspaceSetting(ctx context.Context, delete *DeleteWorkspaceSetting) error
	UpsertWorkspaceSettingV1(ctx context.Context, upsert *storepb.WorkspaceSetting) (*storepb.WorkspaceSetting, error)
	ListWorkspaceSettingsV1(ctx context.Context, find *FindWorkspaceSettingV1) ([]*storepb.WorkspaceSetting, error)

	// User model related methods.
	CreateUser(ctx context.Context, create *User) (*User, error)
	UpdateUser(ctx context.Context, update *UpdateUser) (*User, error)
	ListUsers(ctx context.Context, find *FindUser) ([]*User, error)
	DeleteUser(ctx context.Context, delete *DeleteUser) error

	// UserSetting model related methods.
	UpsertUserSetting(ctx context.Context, upsert *storepb.UserSetting) (*storepb.UserSetting, error)
	ListUserSettings(ctx context.Context, find *FindUserSetting) ([]*storepb.UserSetting, error)

	// IdentityProvider model related methods.
	CreateIdentityProvider(ctx context.Context, create *IdentityProvider) (*IdentityProvider, error)
	ListIdentityProviders(ctx context.Context, find *FindIdentityProvider) ([]*IdentityProvider, error)
	UpdateIdentityProvider(ctx context.Context, update *UpdateIdentityProvider) (*IdentityProvider, error)
	DeleteIdentityProvider(ctx context.Context, delete *DeleteIdentityProvider) error

	// Tag model related methods.
	UpsertTag(ctx context.Context, upsert *Tag) (*Tag, error)
	ListTags(ctx context.Context, find *FindTag) ([]*Tag, error)
	DeleteTag(ctx context.Context, delete *DeleteTag) error

	// Storage model related methods.
	CreateStorage(ctx context.Context, create *Storage) (*Storage, error)
	ListStorages(ctx context.Context, find *FindStorage) ([]*Storage, error)
	UpdateStorage(ctx context.Context, update *UpdateStorage) (*Storage, error)
	DeleteStorage(ctx context.Context, delete *DeleteStorage) error

	// Inbox model related methods.
	CreateInbox(ctx context.Context, create *Inbox) (*Inbox, error)
	ListInboxes(ctx context.Context, find *FindInbox) ([]*Inbox, error)
	UpdateInbox(ctx context.Context, update *UpdateInbox) (*Inbox, error)
	DeleteInbox(ctx context.Context, delete *DeleteInbox) error

	// Webhook model related methods.
	CreateWebhook(ctx context.Context, create *storepb.Webhook) (*storepb.Webhook, error)
	ListWebhooks(ctx context.Context, find *FindWebhook) ([]*storepb.Webhook, error)
	UpdateWebhook(ctx context.Context, update *UpdateWebhook) (*storepb.Webhook, error)
	DeleteWebhook(ctx context.Context, delete *DeleteWebhook) error

	// Reaction model related methods.
	UpsertReaction(ctx context.Context, create *storepb.Reaction) (*storepb.Reaction, error)
	ListReactions(ctx context.Context, find *FindReaction) ([]*storepb.Reaction, error)
	DeleteReaction(ctx context.Context, delete *DeleteReaction) error
}
