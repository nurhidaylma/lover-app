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

func TestLoverService_SetProfile(t *testing.T) {
	initTest(t)
	ctxDummy := context.TODO()

	tests := []struct {
		testName string
		ctx      context.Context
		request  model.Profile

		readUserByIdResp model.User
		readUserByIdErr  error
		writeProfileErr  error

		expectedResp model.ResponseMessage
		expectedErr  error
	}{
		{
			testName: "TC 1: Successful set profile",
			ctx:      ctxDummy,
			request: model.Profile{
				UserId: 1,
				Name:   "Test User",
				Age:    "29",
				Gender: 2,
				Phone:  "082278827291",
			},
			readUserByIdResp: model.User{
				Id: 1,
			},
			readUserByIdErr: nil,
			writeProfileErr: nil,
			expectedResp:    model.ResponseMessage{Message: "success"},
			expectedErr:     nil,
		},
		{
			testName: "TC 2: User is not registered",
			ctx:      ctxDummy,
			request: model.Profile{
				UserId: 1,
				Name:   "Test User",
				Age:    "29",
				Gender: 2,
				Phone:  "082278827291",
			},
			readUserByIdResp: model.User{},
			readUserByIdErr:  nil,
			writeProfileErr:  nil,
			expectedResp:     model.ResponseMessage{Message: "user is not registered"},
			expectedErr:      status.Error(codes.InvalidArgument, codes.InvalidArgument.String()),
		},
		{
			testName: "TC 3: Invalid phone number",
			ctx:      ctxDummy,
			request: model.Profile{
				UserId: 1,
				Name:   "Test User",
				Age:    "29",
				Gender: 2,
				Phone:  "082278",
			},
			expectedResp: model.ResponseMessage{Message: "invalid phone number"},
			expectedErr:  status.Error(codes.InvalidArgument, codes.InvalidArgument.String()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mockDB.EXPECT().ReadUserById(gomock.Any()).Return(tt.readUserByIdResp, tt.readUserByIdErr)
			mockDB.EXPECT().WriteProfile(gomock.Any()).Return(tt.writeProfileErr)

			gotResp, err := services.SetProfile(tt.ctx, tt.request)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Lover.SetProfile() Error = %v, WantError = %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.expectedResp) {
				t.Logf("Lover.SetProfile() Response = %v, SetProfile = %v", gotResp, tt.expectedResp)
				return
			}
		})
	}
}
