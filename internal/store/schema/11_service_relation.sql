-- +migrate Up
CREATE TABLE IF NOT EXISTS instance_at_service
(
    id          BIGSERIAL PRIMARY KEY,
    service_id  BIGINT REFERENCES service (id)  NOT NULL,
    instance_id BIGINT REFERENCES instance (id) NOT NULL,
    created_at  TIMESTAMP                       NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP                       NOT NULL DEFAULT NOW()
);
