package mysql

import (
	"time"

	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/util"
)

const (
	insertUserPurchase = `INSERT INTO user_purchase (user_id, feature_id, purchase_date)
						VALUES(?,?,?)`
)

func (db *dbRepository) WriteUserPurchase(req model.UserPurchase) error {
	const (
		fileName = `user_purchase.go`
		funcName = `WriteUserPurchase`
	)

	timeNow := time.Now().UTC()
	_, err := db.db.Exec(insertUserPurchase, req.UserId, req.FeatureId, timeNow)
	if err != nil {
		return util.NewError(fileName, funcName, err)
	}
	return nil
}
