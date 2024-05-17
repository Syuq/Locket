-- locket_relation
CREATE TABLE locket_relation (
  locket_id INTEGER NOT NULL,
  related_locket_id INTEGER NOT NULL,
  type TEXT NOT NULL,
  UNIQUE(locket_id, related_locket_id, type)
);