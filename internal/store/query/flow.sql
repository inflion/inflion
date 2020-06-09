-- name: GetFlowByName :many
SELECT *
FROM flow
WHERE project_id = $1
  AND flow_name = $2;
