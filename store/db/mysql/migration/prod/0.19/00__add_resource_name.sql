ALTER TABLE `locket` ADD COLUMN `resource_name` VARCHAR(256) AFTER `id`;

UPDATE `locket` SET `resource_name` = uuid();

ALTER TABLE `locket` MODIFY COLUMN `resource_name` VARCHAR(256) NOT NULL;

CREATE UNIQUE INDEX idx_locket_resource_name ON `locket` (`resource_name`);

ALTER TABLE `resource` ADD COLUMN `resource_name` VARCHAR(256) AFTER `id`;

UPDATE `resource` SET `resource_name` = uuid();

ALTER TABLE `resource` MODIFY COLUMN `resource_name` VARCHAR(256) NOT NULL;

CREATE UNIQUE INDEX idx_resource_resource_name ON `resource` (`resource_name`);
