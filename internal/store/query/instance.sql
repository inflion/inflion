-- name: UpsertInstance :one
INSERT INTO instance(instance_id, name, project_id, status)
VALUES ($1, $2, $3, $4)
ON CONFLICT ON CONSTRAINT instance_instance_id_key
    DO UPDATE
    SET name = EXCLUDED.name, status = EXCLUDED.status
RETURNING *;

-- name: SelectInstance :many
SELECT *
FROM instance;

-- name: ResolveIdByInstanceId :one
SELECT id
FROM instance
WHERE
  instance_id = $1
LIMIT 1;

-- name: CreateSecurityGroup :exec
INSERT INTO security_group (security_group_id, security_group_name)
VALUES ($1, $2);