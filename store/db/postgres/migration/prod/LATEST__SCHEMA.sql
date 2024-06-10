-- migration_history
CREATE TABLE migration_history (
  version TEXT NOT NULL PRIMARY KEY,
  created_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
);

-- system_setting
CREATE TABLE system_setting (
  name TEXT NOT NULL PRIMARY KEY,
  value TEXT NOT NULL,
  description TEXT NOT NULL
);

-- user
CREATE TABLE "user" (
  id SERIAL PRIMARY KEY,
  created_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
  updated_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
  row_status TEXT NOT NULL DEFAULT 'NORMAL',
  username TEXT NOT NULL UNIQUE,
  role TEXT NOT NULL DEFAULT 'USER',
  email TEXT NOT NULL DEFAULT '',
  nickname TEXT NOT NULL DEFAULT '',
  password_hash TEXT NOT NULL,
  avatar_url TEXT NOT NULL,
  description TEXT NOT NULL DEFAULT ''
);

-- user_setting
CREATE TABLE user_setting (
  user_id INTEGER NOT NULL,
  key TEXT NOT NULL,
  value TEXT NOT NULL,
  UNIQUE(user_id, key)
);

-- locket
CREATE TABLE locket (
  id SERIAL PRIMARY KEY,
  uid TEXT NOT NULL UNIQUE,
  creator_id INTEGER NOT NULL,
  created_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
  updated_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
  row_status TEXT NOT NULL DEFAULT 'NORMAL',
  content TEXT NOT NULL,
  visibility TEXT NOT NULL DEFAULT 'PRIVATE'
);

-- locket_organizer
CREATE TABLE locket_organizer (
  locket_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  pinned INTEGER NOT NULL DEFAULT 0,
  UNIQUE(locket_id, user_id)
);

-- locket_relation
CREATE TABLE locket_relation (
  locket_id INTEGER NOT NULL,
  related_locket_id INTEGER NOT NULL,
  type TEXT NOT NULL,
  UNIQUE(locket_id, related_locket_id, type)
);

-- resource
CREATE TABLE resource (
  id SERIAL PRIMARY KEY,
  uid TEXT NOT NULL UNIQUE,
  creator_id INTEGER NOT NULL,
  created_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
  updated_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
  filename TEXT NOT NULL,
  blob BYTEA,
  external_link TEXT NOT NULL,
  type TEXT NOT NULL DEFAULT '',
  size INTEGER NOT NULL DEFAULT 0,
  internal_path TEXT NOT NULL DEFAULT '',
  locket_id INTEGER DEFAULT NULL
);

-- tag
CREATE TABLE tag (
  name TEXT NOT NULL,
  creator_id INTEGER NOT NULL,
  UNIQUE(name, creator_id)
);

-- activity
CREATE TABLE activity (
  id SERIAL PRIMARY KEY,
  creator_id INTEGER NOT NULL,
  created_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
  type TEXT NOT NULL DEFAULT '',
  level TEXT NOT NULL DEFAULT 'INFO',
  payload JSONB NOT NULL DEFAULT '{}'
);

-- storage
CREATE TABLE storage (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  type TEXT NOT NULL,
  config JSONB NOT NULL DEFAULT '{}'
);

-- idp
CREATE TABLE idp (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  type TEXT NOT NULL,
  identifier_filter TEXT NOT NULL DEFAULT '',
  config JSONB NOT NULL DEFAULT '{}'
);

-- inbox
CREATE TABLE inbox (
  id SERIAL PRIMARY KEY,
  created_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
  sender_id INTEGER NOT NULL,
  receiver_id INTEGER NOT NULL,
  status TEXT NOT NULL,
  message TEXT NOT NULL
);

-- webhook
CREATE TABLE webhook (
  id SERIAL PRIMARY KEY,
  created_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
  updated_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
  row_status TEXT NOT NULL DEFAULT 'NORMAL',
  creator_id INTEGER NOT NULL,
  name TEXT NOT NULL,
  url TEXT NOT NULL
);

-- reaction
CREATE TABLE reaction (
  id SERIAL PRIMARY KEY,
  created_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
  creator_id INTEGER NOT NULL,
  content_id TEXT NOT NULL,
  reaction_type TEXT NOT NULL,
  UNIQUE(creator_id, content_id, reaction_type)
);
