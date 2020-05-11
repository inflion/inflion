-- +migrate Up
CREATE TABLE IF NOT EXISTS notification_rule
(
    id         BIGSERIAL PRIMARY KEY,
    project_id BIGINT REFERENCES project (id) NOT NULL,
    rule_name  VARCHAR(255)                   NOT NULL,
    rules      JSONB                          NOT NULL DEFAULT jsonb_build_array(),
    created_at TIMESTAMP                      NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP                      NOT NULL DEFAULT NOW()
);