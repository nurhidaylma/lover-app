package service

import (
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nurhidaylma/lover-app.git/internal/repository"
	repoMock "github.com/nurhidaylma/lover-app.git/internal/repository/interfaces"
	"github.com/nurhidaylma/lover-app.git/util"
)

var (
	mockDB    *repoMock.MockDBRepository
	mockRedis *repoMock.MockRedisRepository

	repos    *LoverService
	services Service
)

func initTest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	loggerInstance, err := util.NewCustomLogger("logfile_test.log")
	if err != nil {
		log.Fatal("failed to create logger: ", err.Error())
	}
	util.Logger = loggerInstance

	mockDB = repoMock.NewMockDBRepository(mockCtrl)
	mockRedis = repoMock.NewMockRedisRepository(mockCtrl)

	repos = &LoverService{
		repo: repository.Repository{
			DB:    mockDB,
			Redis: mockRedis,
		},
	}
	services = NewLoverService(
		repository.Repository{
			DB:    mockDB,
			Redis: mockRedis,
		},
	)
}
