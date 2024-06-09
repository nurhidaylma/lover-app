package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nurhidaylma/lover-app.git/internal/model"
	"golang.org/x/crypto/bcrypt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLoverService_Login(t *testing.T) {
	initTest(t)
	dummyPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	tests := []struct {
		testName            string
		email               string
		password            string
		readUserByEmailResp model.User
		readUserByEmailErr  error
		expectedResp        model.User
		expectedErr         error
	}{
		{
			testName:            "TC 1: Valid login",
			email:               "test@example.com",
			password:            "password",
			readUserByEmailResp: model.User{Id: 1, Email: "test@example.com", Password: string(dummyPassword)},
			readUserByEmailErr:  nil,
			expectedResp:        model.User{Id: 1, Email: "test@example.com"},
			expectedErr:         nil,
		},
		{
			testName:            "TC 2: User not found",
			email:               "nonexistent@example.com",
			password:            "password",
			readUserByEmailResp: model.User{Id: 0},
			readUserByEmailErr:  nil,
			expectedResp:        model.User{},
			expectedErr:         status.Error(codes.InvalidArgument, codes.InvalidArgument.String()),
		},
		{
			testName:            "TC 3: Invalid password",
			email:               "test@example.com",
			password:            "wrongpassword",
			readUserByEmailResp: model.User{Id: 1, Email: "test@example.com", Password: string(dummyPassword)},
			readUserByEmailErr:  nil,
			expectedResp:        model.User{},
			expectedErr:         status.Error(codes.InvalidArgument, codes.InvalidArgument.String()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mockDB.EXPECT().ReadUserByEmail(gomock.Any()).Return(tt.readUserByEmailResp, tt.readUserByEmailErr)

			gotResp, err := services.Login(tt.email, tt.password)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Lover.Login() Error = %v, WantError = %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.expectedResp) {
				t.Logf("Lover.Login() Response = %v, Login = %v", gotResp, tt.expectedResp)
				return
			}
		})
	}
}
