-- +migrate Up
CREATE TABLE IF NOT EXISTS project_in_organization
(
    project_id      BIGINT REFERENCES project (id)      NOT NULL,
    organization_id BIGINT REFERENCES organization (id) NOT NULL,
    PRIMARY KEY (project_id, organization_id)
);
-- +migrate Down
DROP TABLE IF EXISTS project_in_organization;