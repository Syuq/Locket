ALTER TABLE locket ADD COLUMN resource_name TEXT NOT NULL DEFAULT "";

UPDATE locket SET resource_name = lower(hex(randomblob(8)));

CREATE UNIQUE INDEX idx_locket_resource_name ON locket (resource_name);

ALTER TABLE resource ADD COLUMN resource_name TEXT NOT NULL DEFAULT "";

UPDATE resource SET resource_name = lower(hex(randomblob(8)));

CREATE UNIQUE INDEX idx_resource_resource_name ON resource (resource_name);
