-- +migrate Up
CREATE TABLE IF NOT EXISTS service
(
    id         BIGSERIAL PRIMARY KEY,
    project_id BIGINT REFERENCES project (id) NOT NULL,
    name       VARCHAR(255)                   NOT NULL,
    created_at TIMESTAMP                      NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP                      NOT NULL DEFAULT NOW()
);
