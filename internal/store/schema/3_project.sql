CREATE TABLE IF NOT EXISTS project
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255)                              NOT NULL,
    description VARCHAR(255),
    created_at  TIMESTAMP                                 NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP                                 NOT NULL DEFAULT NOW(),
    user_id     VARCHAR(255) REFERENCES user_account (id) NOT NULL
);

CREATE TABLE IF NOT EXISTS project_in_organization
(
    project_id      BIGINT REFERENCES project (id)      NOT NULL,
    organization_id BIGINT REFERENCES organization (id) NOT NULL,
    PRIMARY KEY (project_id, organization_id)
);
