// Code generated by sqlc. DO NOT EDIT.

package store

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addToProjectStmt, err = db.PrepareContext(ctx, addToProject); err != nil {
		return nil, fmt.Errorf("error preparing query AddToProject: %w", err)
	}
	if q.allAwsAccountStmt, err = db.PrepareContext(ctx, allAwsAccount); err != nil {
		return nil, fmt.Errorf("error preparing query AllAwsAccount: %w", err)
	}
	if q.confirmInvitationStmt, err = db.PrepareContext(ctx, confirmInvitation); err != nil {
		return nil, fmt.Errorf("error preparing query ConfirmInvitation: %w", err)
	}
	if q.countProjectCollaboratorByUserIdStmt, err = db.PrepareContext(ctx, countProjectCollaboratorByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query CountProjectCollaboratorByUserId: %w", err)
	}
	if q.createSecurityGroupStmt, err = db.PrepareContext(ctx, createSecurityGroup); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSecurityGroup: %w", err)
	}
	if q.getActionsStmt, err = db.PrepareContext(ctx, getActions); err != nil {
		return nil, fmt.Errorf("error preparing query GetActions: %w", err)
	}
	if q.getAwsAccountStmt, err = db.PrepareContext(ctx, getAwsAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetAwsAccount: %w", err)
	}
	if q.getInvitationByTokenStmt, err = db.PrepareContext(ctx, getInvitationByToken); err != nil {
		return nil, fmt.Errorf("error preparing query GetInvitationByToken: %w", err)
	}
	if q.getNotificationRulesStmt, err = db.PrepareContext(ctx, getNotificationRules); err != nil {
		return nil, fmt.Errorf("error preparing query GetNotificationRules: %w", err)
	}
	if q.getSlackWebHooksStmt, err = db.PrepareContext(ctx, getSlackWebHooks); err != nil {
		return nil, fmt.Errorf("error preparing query GetSlackWebHooks: %w", err)
	}
	if q.linkInstanceWithServiceStmt, err = db.PrepareContext(ctx, linkInstanceWithService); err != nil {
		return nil, fmt.Errorf("error preparing query LinkInstanceWithService: %w", err)
	}
	if q.resolveIdByInstanceIdStmt, err = db.PrepareContext(ctx, resolveIdByInstanceId); err != nil {
		return nil, fmt.Errorf("error preparing query ResolveIdByInstanceId: %w", err)
	}
	if q.selectInstanceStmt, err = db.PrepareContext(ctx, selectInstance); err != nil {
		return nil, fmt.Errorf("error preparing query SelectInstance: %w", err)
	}
	if q.upsertInstanceStmt, err = db.PrepareContext(ctx, upsertInstance); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertInstance: %w", err)
	}
	if q.upsertServiceStmt, err = db.PrepareContext(ctx, upsertService); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertService: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addToProjectStmt != nil {
		if cerr := q.addToProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addToProjectStmt: %w", cerr)
		}
	}
	if q.allAwsAccountStmt != nil {
		if cerr := q.allAwsAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing allAwsAccountStmt: %w", cerr)
		}
	}
	if q.confirmInvitationStmt != nil {
		if cerr := q.confirmInvitationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing confirmInvitationStmt: %w", cerr)
		}
	}
	if q.countProjectCollaboratorByUserIdStmt != nil {
		if cerr := q.countProjectCollaboratorByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countProjectCollaboratorByUserIdStmt: %w", cerr)
		}
	}
	if q.createSecurityGroupStmt != nil {
		if cerr := q.createSecurityGroupStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSecurityGroupStmt: %w", cerr)
		}
	}
	if q.getActionsStmt != nil {
		if cerr := q.getActionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getActionsStmt: %w", cerr)
		}
	}
	if q.getAwsAccountStmt != nil {
		if cerr := q.getAwsAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAwsAccountStmt: %w", cerr)
		}
	}
	if q.getInvitationByTokenStmt != nil {
		if cerr := q.getInvitationByTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getInvitationByTokenStmt: %w", cerr)
		}
	}
	if q.getNotificationRulesStmt != nil {
		if cerr := q.getNotificationRulesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNotificationRulesStmt: %w", cerr)
		}
	}
	if q.getSlackWebHooksStmt != nil {
		if cerr := q.getSlackWebHooksStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSlackWebHooksStmt: %w", cerr)
		}
	}
	if q.linkInstanceWithServiceStmt != nil {
		if cerr := q.linkInstanceWithServiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing linkInstanceWithServiceStmt: %w", cerr)
		}
	}
	if q.resolveIdByInstanceIdStmt != nil {
		if cerr := q.resolveIdByInstanceIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing resolveIdByInstanceIdStmt: %w", cerr)
		}
	}
	if q.selectInstanceStmt != nil {
		if cerr := q.selectInstanceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing selectInstanceStmt: %w", cerr)
		}
	}
	if q.upsertInstanceStmt != nil {
		if cerr := q.upsertInstanceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertInstanceStmt: %w", cerr)
		}
	}
	if q.upsertServiceStmt != nil {
		if cerr := q.upsertServiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertServiceStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                   DBTX
	tx                                   *sql.Tx
	addToProjectStmt                     *sql.Stmt
	allAwsAccountStmt                    *sql.Stmt
	confirmInvitationStmt                *sql.Stmt
	countProjectCollaboratorByUserIdStmt *sql.Stmt
	createSecurityGroupStmt              *sql.Stmt
	getActionsStmt                       *sql.Stmt
	getAwsAccountStmt                    *sql.Stmt
	getInvitationByTokenStmt             *sql.Stmt
	getNotificationRulesStmt             *sql.Stmt
	getSlackWebHooksStmt                 *sql.Stmt
	linkInstanceWithServiceStmt          *sql.Stmt
	resolveIdByInstanceIdStmt            *sql.Stmt
	selectInstanceStmt                   *sql.Stmt
	upsertInstanceStmt                   *sql.Stmt
	upsertServiceStmt                    *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                   tx,
		tx:                                   tx,
		addToProjectStmt:                     q.addToProjectStmt,
		allAwsAccountStmt:                    q.allAwsAccountStmt,
		confirmInvitationStmt:                q.confirmInvitationStmt,
		countProjectCollaboratorByUserIdStmt: q.countProjectCollaboratorByUserIdStmt,
		createSecurityGroupStmt:              q.createSecurityGroupStmt,
		getActionsStmt:                       q.getActionsStmt,
		getAwsAccountStmt:                    q.getAwsAccountStmt,
		getInvitationByTokenStmt:             q.getInvitationByTokenStmt,
		getNotificationRulesStmt:             q.getNotificationRulesStmt,
		getSlackWebHooksStmt:                 q.getSlackWebHooksStmt,
		linkInstanceWithServiceStmt:          q.linkInstanceWithServiceStmt,
		resolveIdByInstanceIdStmt:            q.resolveIdByInstanceIdStmt,
		selectInstanceStmt:                   q.selectInstanceStmt,
		upsertInstanceStmt:                   q.upsertInstanceStmt,
		upsertServiceStmt:                    q.upsertServiceStmt,
	}
}
