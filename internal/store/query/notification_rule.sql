-- name: GetNotificationRules :many
SELECT *
FROM notification_rule
WHERE
    project_id = $1;
