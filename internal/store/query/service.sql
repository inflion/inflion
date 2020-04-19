-- name GetServices :many
SELECT * FROM service;

-- name: UpsertService :one
INSERT INTO service(name, project_id)
VALUES ($1, $2)
ON CONFLICT ON CONSTRAINT service_project_id_name_key
    DO UPDATE
    SET name = EXCLUDED.name
RETURNING *;

-- name: LinkInstanceWithService :one
INSERT INTO instance_at_service(service_id, instance_id)
VALUES ($1, $2)
ON CONFLICT ON CONSTRAINT instance_at_service_service_id_instance_id_key
  DO NOTHING
RETURNING *;
