// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: resume_points.sql

package pgmodel

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createResumePoint = `-- name: CreateResumePoint :one
INSERT INTO resume_points (
    id, 
    type,
    app_id, 
    command_id, 
    event_listener_id, 
    message_id, 
    message_instance_id,
    flow_source_id, 
    flow_node_id, 
    flow_state, 
    created_at, 
    expires_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id, type, app_id, command_id, event_listener_id, message_id, message_instance_id, flow_source_id, flow_node_id, flow_state, created_at, expires_at
`

type CreateResumePointParams struct {
	ID                string
	Type              string
	AppID             string
	CommandID         pgtype.Text
	EventListenerID   pgtype.Text
	MessageID         pgtype.Text
	MessageInstanceID pgtype.Int8
	FlowSourceID      pgtype.Text
	FlowNodeID        string
	FlowState         []byte
	CreatedAt         pgtype.Timestamp
	ExpiresAt         pgtype.Timestamp
}

func (q *Queries) CreateResumePoint(ctx context.Context, arg CreateResumePointParams) (ResumePoint, error) {
	row := q.db.QueryRow(ctx, createResumePoint,
		arg.ID,
		arg.Type,
		arg.AppID,
		arg.CommandID,
		arg.EventListenerID,
		arg.MessageID,
		arg.MessageInstanceID,
		arg.FlowSourceID,
		arg.FlowNodeID,
		arg.FlowState,
		arg.CreatedAt,
		arg.ExpiresAt,
	)
	var i ResumePoint
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.AppID,
		&i.CommandID,
		&i.EventListenerID,
		&i.MessageID,
		&i.MessageInstanceID,
		&i.FlowSourceID,
		&i.FlowNodeID,
		&i.FlowState,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const deleteExpiredResumePoints = `-- name: DeleteExpiredResumePoints :exec
DELETE FROM resume_points WHERE expires_at < $1
`

func (q *Queries) DeleteExpiredResumePoints(ctx context.Context, expiresAt pgtype.Timestamp) error {
	_, err := q.db.Exec(ctx, deleteExpiredResumePoints, expiresAt)
	return err
}

const deleteResumePoint = `-- name: DeleteResumePoint :exec
DELETE FROM resume_points WHERE id = $1
`

func (q *Queries) DeleteResumePoint(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteResumePoint, id)
	return err
}

const resumePoint = `-- name: ResumePoint :one
SELECT id, type, app_id, command_id, event_listener_id, message_id, message_instance_id, flow_source_id, flow_node_id, flow_state, created_at, expires_at FROM resume_points WHERE id = $1
`

func (q *Queries) ResumePoint(ctx context.Context, id string) (ResumePoint, error) {
	row := q.db.QueryRow(ctx, resumePoint, id)
	var i ResumePoint
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.AppID,
		&i.CommandID,
		&i.EventListenerID,
		&i.MessageID,
		&i.MessageInstanceID,
		&i.FlowSourceID,
		&i.FlowNodeID,
		&i.FlowState,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}
