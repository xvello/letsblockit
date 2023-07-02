// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: qUsers.sql

package db

import (
	"context"
	"time"
)

const addUserBan = `-- name: AddUserBan :exec
INSERT INTO banned_users (user_id, reason)
VALUES ($1, $2)
`

type AddUserBanParams struct {
	UserID string
	Reason string
}

func (q *Queries) AddUserBan(ctx context.Context, arg AddUserBanParams) error {
	_, err := q.db.Exec(ctx, addUserBan, arg.UserID, arg.Reason)
	return err
}

const getBannedUsers = `-- name: GetBannedUsers :many
SELECT user_id
from banned_users
WHERE lifted_at IS NULL
`

func (q *Queries) GetBannedUsers(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, getBannedUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var user_id string
		if err := rows.Scan(&user_id); err != nil {
			return nil, err
		}
		items = append(items, user_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserPreferences = `-- name: GetUserPreferences :one
SELECT user_id, news_cursor, beta_features, color_mode
FROM user_preferences
WHERE user_id = $1
`

func (q *Queries) GetUserPreferences(ctx context.Context, userID string) (UserPreference, error) {
	row := q.db.QueryRow(ctx, getUserPreferences, userID)
	var i UserPreference
	err := row.Scan(
		&i.UserID,
		&i.NewsCursor,
		&i.BetaFeatures,
		&i.ColorMode,
	)
	return i, err
}

const initUserPreferences = `-- name: InitUserPreferences :one
INSERT INTO user_preferences (user_id)
VALUES ($1)
RETURNING user_id, news_cursor, beta_features, color_mode
`

func (q *Queries) InitUserPreferences(ctx context.Context, userID string) (UserPreference, error) {
	row := q.db.QueryRow(ctx, initUserPreferences, userID)
	var i UserPreference
	err := row.Scan(
		&i.UserID,
		&i.NewsCursor,
		&i.BetaFeatures,
		&i.ColorMode,
	)
	return i, err
}

const liftUserBan = `-- name: LiftUserBan :exec
UPDATE banned_users
SET lift_reason = $2::text,
    lifted_at   = NOW()
WHERE (user_id = $1)
`

type LiftUserBanParams struct {
	UserID string
	Reason string
}

func (q *Queries) LiftUserBan(ctx context.Context, arg LiftUserBanParams) error {
	_, err := q.db.Exec(ctx, liftUserBan, arg.UserID, arg.Reason)
	return err
}

const updateNewsCursor = `-- name: UpdateNewsCursor :exec
UPDATE user_preferences
SET news_cursor = $2
WHERE user_id = $1
`

type UpdateNewsCursorParams struct {
	UserID     string
	NewsCursor time.Time
}

func (q *Queries) UpdateNewsCursor(ctx context.Context, arg UpdateNewsCursorParams) error {
	_, err := q.db.Exec(ctx, updateNewsCursor, arg.UserID, arg.NewsCursor)
	return err
}

const updateUserPreferences = `-- name: UpdateUserPreferences :exec
UPDATE user_preferences
SET color_mode    = $2,
    beta_features = $3
WHERE user_id = $1
`

type UpdateUserPreferencesParams struct {
	UserID       string
	ColorMode    ColorMode
	BetaFeatures bool
}

func (q *Queries) UpdateUserPreferences(ctx context.Context, arg UpdateUserPreferencesParams) error {
	_, err := q.db.Exec(ctx, updateUserPreferences, arg.UserID, arg.ColorMode, arg.BetaFeatures)
	return err
}
