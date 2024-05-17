-- change locket visibility field from "PRIVATE"/"PUBLIC" to "PRIVATE"/"PROTECTED"/"PUBLIC".
PRAGMA foreign_keys = off;

DROP TABLE IF EXISTS _locket_old;

ALTER TABLE
  locket RENAME TO _locket_old;

CREATE TABLE locket (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  creator_id INTEGER NOT NULL,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  row_status TEXT NOT NULL CHECK (row_status IN ('NORMAL', 'ARCHIVED')) DEFAULT 'NORMAL',
  content TEXT NOT NULL DEFAULT '',
  visibility TEXT NOT NULL CHECK (visibility IN ('PUBLIC', 'PROTECTED', 'PRIVATE')) DEFAULT 'PRIVATE',
  FOREIGN KEY(creator_id) REFERENCES user(id) ON DELETE CASCADE
);

INSERT INTO
  locket (
    id,
    creator_id,
    created_ts,
    updated_ts,
    row_status,
    content,
    visibility
  )
SELECT
  id,
  creator_id,
  created_ts,
  updated_ts,
  row_status,
  content,
  visibility
FROM
  _locket_old;

DROP TABLE IF EXISTS _locket_old;

PRAGMA foreign_keys = on;