-- name: GetSlackWebHooks :many
SELECT *
FROM slack_webhook
WHERE project_id = $1;
