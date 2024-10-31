// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 0002_pre_go_acc_user_base_9999.sql

package database

import (
	"context"
	"database/sql"
)

const addUserBase = `-- name: AddUserBase :execresult
INSERT INTO pre_go_acc_user_base_9999 (
       user_account, user_password, user_salt, user_created_at, user_updated_at
) VALUES (
      ?, ?, ?, NOW(), NOW()
     )
`

type AddUserBaseParams struct {
	UserAccount  string
	UserPassword string
	UserSalt     string
}

func (q *Queries) AddUserBase(ctx context.Context, arg AddUserBaseParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addUserBase, arg.UserAccount, arg.UserPassword, arg.UserSalt)
}

const checkUserHasExists = `-- name: CheckUserHasExists :one
SELECT COUNT(*)
FROM ` + "`" + `pre_go_acc_user_base_9999` + "`" + `
WHERE user_account = ?
`

func (q *Queries) CheckUserHasExists(ctx context.Context, userAccount string) (int64, error) {
	row := q.db.QueryRowContext(ctx, checkUserHasExists, userAccount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getOneUserInfo = `-- name: GetOneUserInfo :one
SELECT user_id, user_account, user_password, user_salt
FROM ` + "`" + `pre_go_acc_user_base_9999` + "`" + `
WHERE user_account = ?
`

type GetOneUserInfoRow struct {
	UserID       int32
	UserAccount  string
	UserPassword string
	UserSalt     string
}

func (q *Queries) GetOneUserInfo(ctx context.Context, userAccount string) (GetOneUserInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getOneUserInfo, userAccount)
	var i GetOneUserInfoRow
	err := row.Scan(
		&i.UserID,
		&i.UserAccount,
		&i.UserPassword,
		&i.UserSalt,
	)
	return i, err
}

const getOneUserInfoAdmin = `-- name: GetOneUserInfoAdmin :one
SELECT user_id, user_account, user_password, user_salt, user_login_time,
        user_logout_time, user_created_at, user_updated_at
FROM ` + "`" + `pre_go_acc_user_base_9999` + "`" + `
WHERE user_account = ?
`

type GetOneUserInfoAdminRow struct {
	UserID         int32
	UserAccount    string
	UserPassword   string
	UserSalt       string
	UserLoginTime  sql.NullTime
	UserLogoutTime sql.NullTime
	UserCreatedAt  sql.NullTime
	UserUpdatedAt  sql.NullTime
}

func (q *Queries) GetOneUserInfoAdmin(ctx context.Context, userAccount string) (GetOneUserInfoAdminRow, error) {
	row := q.db.QueryRowContext(ctx, getOneUserInfoAdmin, userAccount)
	var i GetOneUserInfoAdminRow
	err := row.Scan(
		&i.UserID,
		&i.UserAccount,
		&i.UserPassword,
		&i.UserSalt,
		&i.UserLoginTime,
		&i.UserLogoutTime,
		&i.UserCreatedAt,
		&i.UserUpdatedAt,
	)
	return i, err
}