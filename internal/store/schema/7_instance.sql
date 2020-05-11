-- +migrate Up
CREATE TABLE instance
(
    id          BIGSERIAL PRIMARY KEY,
    name        character varying(255)            NOT NULL,
    instance_id character varying(255)            NOT NULL,
    status      VARCHAR(255)                      NOT NULL DEFAULT 'unknown',
    created_at  TIMESTAMP                         NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP                         NOT NULL DEFAULT NOW(),
    project_id  BIGSERIAL REFERENCES project (id) NOT NULL,
    UNIQUE (instance_id)
);