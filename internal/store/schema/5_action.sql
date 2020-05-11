-- +migrate Up
CREATE TABLE IF NOT EXISTS action
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    VARCHAR(255) REFERENCES user_account (id) NOT NULL,
    project_id BIGINT REFERENCES project (id)            NOT NULL,
    name       VARCHAR(255)                              NOT NULL,
    body       JSONB                                     NOT NULL DEFAULT jsonb_build_array(),
    created_at TIMESTAMP                                 NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP                                 NOT NULL DEFAULT NOW()
);