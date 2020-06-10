-- +migrate Up
CREATE TABLE flow
(
    id         bigint                                                   NOT NULL,
    project_id bigint                                                   NOT NULL,
    flow_name  character varying(255)                                   NOT NULL,
    body       jsonb                       DEFAULT jsonb_build_object() NOT NULL,
    created_at timestamp without time zone DEFAULT now()                NOT NULL,
    updated_at timestamp without time zone DEFAULT now()                NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS flow;
