-- name: GetActions :many
SELECT *
FROM action
WHERE project_id = $1;
