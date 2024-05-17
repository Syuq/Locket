DROP TABLE IF EXISTS locket_organizer_temp;

CREATE TABLE locket_organizer_temp (
  locket_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  pinned INTEGER NOT NULL CHECK (pinned IN (0, 1)) DEFAULT 0,
  UNIQUE(locket_id, user_id)
);

INSERT INTO
  locket_organizer_temp (locket_id, user_id, pinned)
SELECT
  locket_id,
  user_id,
  pinned
FROM
  locket_organizer;

DROP TABLE locket_organizer;

ALTER TABLE
  locket_organizer_temp RENAME TO locket_organizer;