-- name: GetAwsAccount :one
SELECT account_id, role_name, external_id, project_id
FROM aws_account
WHERE
    project_id = $1
LIMIT 1;

-- name: AllAwsAccount :many
SELECT account_id, role_name, external_id, project_id
FROM aws_account;
