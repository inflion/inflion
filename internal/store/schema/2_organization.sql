-- +migrate Up
CREATE TABLE organization (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	unique_name VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +migrate Down
DROP TABLE IF EXISTS organization;