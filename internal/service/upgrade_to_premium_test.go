package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/nurhidaylma/lover-app.git/internal/model"
)

func TestLoverService_UpgradeToPremium(t *testing.T) {
	initTest(t)
	ctxDummy := context.TODO()

	tests := []struct {
		testName string
		ctx      context.Context
		request  model.UserPurchase

		writeUserPurchaseErr error
		setPremiumUserErr    error

		expectedResp model.ResponseMessage
		expectedErr  error
	}{
		{
			testName: "TC 1: Successful upgrade to premium",
			ctx:      ctxDummy,
			request: model.UserPurchase{
				UserId:       1,
				FeatureId:    1,
				PurchaseDate: time.Now(),
			},
			writeUserPurchaseErr: nil,
			setPremiumUserErr:    nil,
			expectedResp:         model.ResponseMessage{Message: "success"},
			expectedErr:          nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mockDB.EXPECT().WriteUserPurchase(gomock.Any()).Return(tt.writeUserPurchaseErr)
			mockRedis.EXPECT().SetPremiumUser(gomock.Any()).Return(tt.setPremiumUserErr)

			gotResp, err := services.UpgradeToPremium(tt.ctx, tt.request)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Lover.UpgradeToPremium() Error = %v, WantError = %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.expectedResp) {
				t.Logf("Lover.UpgradeToPremium() Response = %v, UpgradeToPremium = %v", gotResp, tt.expectedResp)
				return
			}
		})
	}
}
