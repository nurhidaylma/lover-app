package mysql

import (
	"database/sql"
	"time"

	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/util"
)

const (
	insertUser = `INSERT INTO users (email, username, password, premium_status, created_at, updated_at)
				VALUES (?,?,?,?,?,?)`
	selectUserById       = `SELECT email, username, premium_status FROM users WHERE id = ?`
	selectUserByEmail    = `SELECT email, username, premium_status FROM users WHERE email = ?`
	selectUserByUserName = `SELECT email, username, premium_status FROM users WHERE username = ?`
)

func (db *dbRepository) WriteUser(req model.User) error {
	const (
		fileName = `user.go`
		funcName = "WriteUser"
	)

	timeNow := time.Now().UTC()
	_, err := db.db.Exec(insertUser, req.Email, req.UserName, util.IsNotPremium, timeNow, timeNow)
	if err != nil {
		return util.NewError(fileName, funcName, err)
	}
	return nil
}

func (db *dbRepository) ReadUserById(id int) (resp model.User, err error) {
	const (
		fileName = `user.go`
		funcName = "ReadUserById"
	)

	err = db.db.QueryRow(selectUserById, id).Scan(
		&resp.Email,
		&resp.UserName,
		&resp.PremiumStatus,
	)
	if err != nil && err != sql.ErrNoRows {
		return resp, util.NewError(fileName, funcName, err)
	}

	return
}

func (db *dbRepository) ReadUserByEmail(email string) (resp model.User, err error) {
	const (
		fileName = `user.go`
		funcName = "ReadUserByEmail"
	)

	err = db.db.QueryRow(selectUserByEmail, email).Scan(
		&resp.Email,
		&resp.UserName,
		&resp.PremiumStatus,
	)
	if err != nil && err != sql.ErrNoRows {
		return resp, util.NewError(fileName, funcName, err)
	}

	return
}

func (db *dbRepository) ReadUserByUserName(username string) (resp model.User, err error) {
	const (
		fileName = `user.go`
		funcName = "ReadUserByUserName"
	)

	err = db.db.QueryRow(selectUserByEmail, username).Scan(
		&resp.Email,
		&resp.UserName,
		&resp.PremiumStatus,
	)
	if err != nil && err != sql.ErrNoRows {
		return resp, util.NewError(fileName, funcName, err)
	}

	return
}
