ALTER TABLE locket ADD COLUMN resource_name TEXT;

UPDATE locket SET resource_name = uuid_in(md5(random()::text || random()::text)::cstring);

ALTER TABLE locket ALTER COLUMN resource_name SET NOT NULL;

CREATE UNIQUE INDEX idx_locket_resource_name ON locket (resource_name);

ALTER TABLE resource ADD COLUMN resource_name TEXT;

UPDATE resource SET resource_name = uuid_in(md5(random()::text || random()::text)::cstring);

ALTER TABLE resource ALTER COLUMN resource_name SET NOT NULL;

CREATE UNIQUE INDEX idx_resource_resource_name ON resource (resource_name);
