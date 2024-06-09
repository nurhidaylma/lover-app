package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nurhidaylma/lover-app.git/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLoverService_Signup(t *testing.T) {
	initTest(t)

	tests := []struct {
		testName string
		request  model.User

		readUserByEmailResp    model.User
		readUserByEmailErr     error
		readUserByUserNameResp model.User
		readUserByUserNameErr  error
		writeUserErr           error

		expectedResp model.ResponseMessage
		expectedErr  error
	}{
		{
			testName: "TC 1: Successful signup",
			request: model.User{
				Email:    "valid@example.com",
				UserName: "validuser",
				Password: "validpassword",
			},
			readUserByEmailResp:    model.User{},
			readUserByEmailErr:     nil,
			readUserByUserNameResp: model.User{},
			readUserByUserNameErr:  nil,
			writeUserErr:           nil,
			expectedResp:           model.ResponseMessage{Message: "success"},
			expectedErr:            nil,
		},
		{
			testName: "TC 2: Invalid email address",
			request: model.User{
				Email:    "invalid-email",
				UserName: "validuser",
				Password: "validpassword",
			},
			expectedResp: model.ResponseMessage{Message: "invalid email address"},
			expectedErr:  status.Error(codes.InvalidArgument, codes.InvalidArgument.String()),
		},
		{
			testName: "TC 3: Email already registered",
			request: model.User{
				Email:    "registered@example.com",
				UserName: "validuser",
				Password: "validpassword",
			},
			readUserByEmailResp: model.User{
				Id:    1,
				Email: "registered@example.com",
			},
			expectedResp: model.ResponseMessage{Message: "user already registered"},
			expectedErr:  status.Error(codes.InvalidArgument, codes.InvalidArgument.String()),
		},
		{
			testName: "TC 4: Username already registered",
			request: model.User{
				Email:    "registered@example.com",
				UserName: "validuser",
				Password: "validpassword",
			},
			readUserByEmailResp: model.User{
				Id:       1,
				UserName: "validuser",
			},
			expectedResp: model.ResponseMessage{Message: "user already registered"},
			expectedErr:  status.Error(codes.InvalidArgument, codes.InvalidArgument.String()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mockDB.EXPECT().ReadUserByEmail(gomock.Any()).Return(tt.readUserByEmailResp, tt.readUserByEmailErr)
			mockDB.EXPECT().ReadUserByUserName(gomock.Any()).Return(tt.readUserByUserNameResp, tt.readUserByUserNameErr)
			mockDB.EXPECT().WriteUser(gomock.Any()).Return(tt.writeUserErr)

			gotResp, err := services.SignUp(tt.request)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Lover.SignUp() Error = %v, WantError = %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.expectedResp) {
				t.Logf("Lover.SignUp() Response = %v, SignUp = %v", gotResp, tt.expectedResp)
				return
			}
		})
	}
}
