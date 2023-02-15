// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddUserBan(ctx context.Context, arg AddUserBanParams) error
	CountInstances(ctx context.Context, arg CountInstancesParams) (int64, error)
	CountListsForUser(ctx context.Context, userID string) (int64, error)
	CreateInstance(ctx context.Context, arg CreateInstanceParams) error
	CreateListForUser(ctx context.Context, userID string) (uuid.UUID, error)
	DeleteInstance(ctx context.Context, arg DeleteInstanceParams) error
	GetBannedUsers(ctx context.Context) ([]string, error)
	GetInstance(ctx context.Context, arg GetInstanceParams) (GetInstanceRow, error)
	GetInstanceStats(ctx context.Context) ([]GetInstanceStatsRow, error)
	GetInstancesForList(ctx context.Context, listID int32) ([]GetInstancesForListRow, error)
	GetInstancesForUser(ctx context.Context, userID string) ([]GetInstancesForUserRow, error)
	GetListForToken(ctx context.Context, token uuid.UUID) (GetListForTokenRow, error)
	GetListForUser(ctx context.Context, userID string) (GetListForUserRow, error)
	GetStats(ctx context.Context) (GetStatsRow, error)
	GetUserPreferences(ctx context.Context, userID string) (UserPreference, error)
	InitUserPreferences(ctx context.Context, userID string) (UserPreference, error)
	LiftUserBan(ctx context.Context, arg LiftUserBanParams) error
	MarkListDownloaded(ctx context.Context, token uuid.UUID) error
	RotateListToken(ctx context.Context, arg RotateListTokenParams) error
	UpdateInstance(ctx context.Context, arg UpdateInstanceParams) error
	UpdateNewsCursor(ctx context.Context, arg UpdateNewsCursorParams) error
}

var _ Querier = (*Queries)(nil)
