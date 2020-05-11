-- +migrate Up
CREATE TABLE IF NOT EXISTS aws_account
(
    id          BIGSERIAL PRIMARY KEY,
    user_id     VARCHAR(255) REFERENCES user_account (id) NOT NULL,
    project_id  BIGINT REFERENCES project (id)            NOT NULL,
    account_id  VARCHAR(255)                              NOT NULL,
    role_name   VARCHAR(255)                              NOT NULL,
    external_id VARCHAR(255)                              NOT NULL,
    created_at  TIMESTAMP                                 NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP                                 NOT NULL DEFAULT NOW()
);
