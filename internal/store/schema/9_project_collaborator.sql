-- +migrate Up
CREATE TABLE IF NOT EXISTS project_collaborator
(
    user_id    VARCHAR(255) REFERENCES user_account (id) NOT NULL,
    project_id BIGINT REFERENCES project (id)            NOT NULL,
    created_at TIMESTAMP                                 NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP                                 NOT NULL DEFAULT NOW()
);
