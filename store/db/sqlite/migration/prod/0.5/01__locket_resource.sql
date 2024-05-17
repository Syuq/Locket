-- locket_resource
CREATE TABLE locket_resource (
  locket_id INTEGER NOT NULL,
  resource_id INTEGER NOT NULL,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  FOREIGN KEY(locket_id) REFERENCES locket(id) ON DELETE CASCADE,
  FOREIGN KEY(resource_id) REFERENCES resource(id) ON DELETE CASCADE,
  UNIQUE(locket_id, resource_id)
);