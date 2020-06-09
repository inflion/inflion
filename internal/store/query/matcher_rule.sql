-- name: GetMatcherRules :many
SELECT *
FROM matcher_rule
WHERE
    project_id = $1;
