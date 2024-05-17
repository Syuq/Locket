ALTER TABLE resource ADD COLUMN locket_id INTEGER;

UPDATE resource
SET locket_id = (
  SELECT locket_id
  FROM locket_resource
  WHERE resource.id = locket_resource.resource_id
  LIMIT 1
);

CREATE INDEX idx_resource_locket_id ON resource (locket_id);

DROP TABLE IF EXISTS locket_resource;
