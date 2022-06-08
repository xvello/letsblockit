// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: qUsers.sql

package db

import (
	"context"
	"time"
)

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
SELECT user_id, news_cursor
FROM user_preferences
WHERE user_id = $1
`

func (q *Queries) GetUserPreferences(ctx context.Context, userID string) (UserPreference, error) {
	row := q.db.QueryRow(ctx, getUserPreferences, userID)
	var i UserPreference
	err := row.Scan(&i.UserID, &i.NewsCursor)
	return i, err
}

const initUserPreferences = `-- name: InitUserPreferences :one
INSERT INTO user_preferences (user_id)
VALUES ($1)
RETURNING user_id, news_cursor
`

func (q *Queries) InitUserPreferences(ctx context.Context, userID string) (UserPreference, error) {
	row := q.db.QueryRow(ctx, initUserPreferences, userID)
	var i UserPreference
	err := row.Scan(&i.UserID, &i.NewsCursor)
	return i, err
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
