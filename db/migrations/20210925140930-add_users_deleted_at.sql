
-- +migrate Up
ALTER TABLE users ADD deleted_at DATETIME DEFAULT NULL;
-- +migrate Down
ALTER TABLE user DROP COLUMN deleted_at;