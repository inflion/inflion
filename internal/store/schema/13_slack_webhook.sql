-- +migrate Up
CREATE TABLE IF NOT EXISTS slack_webhook
(
    id          BIGSERIAL PRIMARY KEY,
    user_id     character varying(255)      NOT NULL REFERENCES user_account (id),
    project_id  bigint                      NOT NULL REFERENCES project (id),
    name        character varying           NOT NULL,
    channel     character varying           NOT NULL,
    webhook_url character varying(255)      NOT NULL,
    created_at  timestamp without time zone NOT NULL DEFAULT now(),
    updated_at  timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT slack_webhook_project_id_channel_key UNIQUE (project_id, channel)
);
