package mysql

import (
	"time"

	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/util"
)

const (
	insertProfile = `INSERT INTO profiles (user_id, name, age, gender, phone, created_at, updated_at) 
					VALUES (?,?,?,?,?,?,?)`
)

func (db *dbRepository) WriteProfile(req model.Profile) error {
	const (
		fileName = `profile.go`
		funcName = `WriteProfile`
	)

	timeNow := time.Now().UTC()
	_, err := db.db.Exec(insertProfile,
		req.UserId,
		req.Name,
		req.Age,
		req.Gender,
		req.Phone,
		timeNow,
		timeNow)
	if err != nil {
		return util.NewError(fileName, funcName, err)
	}

	return nil
}
