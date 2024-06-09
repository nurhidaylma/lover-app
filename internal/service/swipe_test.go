package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nurhidaylma/lover-app.git/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLoverService_Swipe(t *testing.T) {
	initTest(t)
	ctxDummy := context.TODO()

	tests := []struct {
		testName string
		request  model.Swipe
		ctx      context.Context

		getSwipeCountResp  int
		getSwipeCountErr   error
		getPremiumUserResp bool
		getPremiumUserErr  error
		writeSwipeErr      error
		setSwipeCountErr   error

		expectedResp model.ResponseMessage
		expectedErr  error
	}{
		{
			testName: "TC 1: Successful swipe for non-premium user",
			request: model.Swipe{
				UserId:    1,
				ProfileId: 10,
				SwipeType: 1,
			},
			ctx:                ctxDummy,
			getSwipeCountResp:  3,
			getSwipeCountErr:   nil,
			getPremiumUserResp: false,
			getPremiumUserErr:  nil,
			writeSwipeErr:      nil,
			setSwipeCountErr:   nil,
			expectedResp:       model.ResponseMessage{Message: "success"},
			expectedErr:        nil,
		},
		{
			testName: "TC 2: Successful swipe for premium user",
			request: model.Swipe{
				UserId:    1,
				ProfileId: 10,
				SwipeType: 1,
			},
			ctx:                ctxDummy,
			getSwipeCountResp:  20,
			getSwipeCountErr:   nil,
			getPremiumUserResp: true,
			getPremiumUserErr:  nil,
			writeSwipeErr:      nil,
			setSwipeCountErr:   nil,
			expectedResp:       model.ResponseMessage{Message: "success"},
			expectedErr:        nil,
		},
		{
			testName: "TC 3: Invalid swipe type",
			request: model.Swipe{
				UserId:    1,
				ProfileId: 10,
				SwipeType: 5,
			},
			ctx:          ctxDummy,
			expectedResp: model.ResponseMessage{Message: "invalid swipe type"},
			expectedErr:  status.Error(codes.InvalidArgument, codes.InvalidArgument.String()),
		},
		{
			testName: "TC 4: Swipe reaches limit",
			request: model.Swipe{
				UserId:    1,
				ProfileId: 10,
				SwipeType: -1,
			},
			ctx:                ctxDummy,
			getSwipeCountResp:  20,
			getSwipeCountErr:   nil,
			getPremiumUserResp: false,
			getPremiumUserErr:  nil,
			expectedResp:       model.ResponseMessage{Message: "swipe reached limit"},
			expectedErr:        status.Error(codes.InvalidArgument, codes.InvalidArgument.String()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mockRedis.EXPECT().GetSwipeCount(gomock.Any()).Return(tt.getSwipeCountResp, tt.getSwipeCountErr)
			mockRedis.EXPECT().GetPremiumUser(gomock.Any()).Return(tt.getPremiumUserResp, tt.getPremiumUserErr)
			mockDB.EXPECT().WriteSwipe(gomock.Any()).Return(tt.writeSwipeErr)
			mockRedis.EXPECT().SetSwipeCount(gomock.Any()).Return(tt.setSwipeCountErr)

			gotResp, err := services.Swipe(tt.ctx, tt.request)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Lover.Swipe() Error = %v, WantError = %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.expectedResp) {
				t.Logf("Lover.Swipe() Response = %v, Swipe = %v", gotResp, tt.expectedResp)
				return
			}
		})
	}
}
