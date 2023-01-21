// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type BannedUser struct {
	ID         int32
	UserID     string
	CreatedAt  time.Time
	Reason     string
	LiftedAt   sql.NullTime
	LiftReason sql.NullString
}

type FilterInstance struct {
	ID           int32
	UserID       string
	ListID       int32
	TemplateName string
	Params       pgtype.JSONB
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
	TestMode     bool
}

type FilterList struct {
	ID           int32
	UserID       string
	Token        uuid.UUID
	CreatedAt    time.Time
	Downloaded   bool
	DownloadedAt sql.NullTime
}

type UserPreference struct {
	UserID       string
	NewsCursor   time.Time
	BetaFeatures bool
}
