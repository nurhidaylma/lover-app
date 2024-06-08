package mysql

import (
	"time"

	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/util"
)

const (
	insertSwipe = `INSERT INTO swipe (user_id, profile_id, swipe_type, created_at)`
)

func (db *dbRepository) WriteSwipe(req model.Swipe) error {
	const (
		fileName = `swipe.go`
		funcName = `WriteSwipe`
	)

	timeNow := time.Now().UTC()
	_, err := db.db.Exec(insertSwipe, req.UserId, req.ProfileId, req.SwipeType, timeNow)
	if err != nil {
		return util.NewError(fileName, funcName, err)
	}
	return nil
}
