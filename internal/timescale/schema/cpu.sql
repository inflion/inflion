CREATE TABLE metrics
(
    time        TIMESTAMPTZ      NOT NULL,
    instance_id VARCHAR(255)     NOT NULL,
    type        VARCHAR(255)     NOT NULL,
    value       DOUBLE PRECISION NULL
);
