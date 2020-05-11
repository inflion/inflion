-- +migrate Up
CREATE TABLE IF NOT EXISTS project_invitation
(
    id              BIGSERIAL PRIMARY KEY,
    invitee_user_id VARCHAR(255) REFERENCES user_account (id) NOT NULL,
    project_id      BIGINT REFERENCES project (id)            NOT NULL,
    mail_address    VARCHAR(255)                              NOT NULL,
    token           VARCHAR(255)                              NOT NULL DEFAULT rand(),
    confirmed       BOOLEAN                                   NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMP                                 NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP                                 NOT NULL DEFAULT NOW()
);