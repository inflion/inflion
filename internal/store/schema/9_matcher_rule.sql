-- +migrate Up
CREATE TABLE matcher_rule
(
    id               bigint                                                  NOT NULL,
    project_id       bigint REFERENCES project (id)                          NOT NULL,
    rule_name        character varying(255)                                  NOT NULL,
    target_flow_name character varying(255)                                  NOT NULL,
    rules            jsonb                       DEFAULT jsonb_build_array() NOT NULL,
    created_at       timestamp without time zone DEFAULT now()               NOT NULL,
    updated_at       timestamp without time zone DEFAULT now()               NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS matcher_rule;