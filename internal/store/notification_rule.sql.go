// Code generated by sqlc. DO NOT EDIT.
// source: notification_rule.sql

package store

import (
	"context"
)

const getNotificationRules = `-- name: GetNotificationRules :many
SELECT id, project_id, rule_name, rules, created_at, updated_at
FROM notification_rule
WHERE
    project_id = $1
`

func (q *Queries) GetNotificationRules(ctx context.Context, projectID int64) ([]NotificationRule, error) {
	rows, err := q.query(ctx, q.getNotificationRulesStmt, getNotificationRules, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []NotificationRule
	for rows.Next() {
		var i NotificationRule
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.RuleName,
			&i.Rules,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
