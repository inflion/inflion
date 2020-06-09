// Code generated by sqlc. DO NOT EDIT.
// source: 15_flow.sql

package store

import (
	"context"
)

const getFlowByName = `-- name: GetFlowByName :many
SELECT id, project_id, flow_name, body, created_at, updated_at
FROM flow
WHERE project_id = $1
  AND flow_name = $2
`

type GetFlowByNameParams struct {
	ProjectID int64
	FlowName  string
}

func (q *Queries) GetFlowByName(ctx context.Context, arg GetFlowByNameParams) ([]Flow, error) {
	rows, err := q.query(ctx, q.getFlowByNameStmt, getFlowByName, arg.ProjectID, arg.FlowName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Flow
	for rows.Next() {
		var i Flow
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.FlowName,
			&i.Body,
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
