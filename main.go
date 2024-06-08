package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nurhidaylma/lover-app.git/config"
	"github.com/nurhidaylma/lover-app.git/internal/endpoint"
	"github.com/nurhidaylma/lover-app.git/internal/repository"
	"github.com/nurhidaylma/lover-app.git/internal/repository/mysql"
	"github.com/nurhidaylma/lover-app.git/internal/repository/redis"
	"github.com/nurhidaylma/lover-app.git/internal/service"
	"github.com/nurhidaylma/lover-app.git/internal/transport/rest"
	"github.com/nurhidaylma/lover-app.git/util"
)

func main() {
	loggerInstance, err := util.NewCustomLogger("logfile.log")
	if err != nil {
		log.Fatal("failed to create logger: ", err.Error())
	}
	util.Logger = loggerInstance

	var config config.ServiceConfig
	path, err := os.Getwd()
	if err != nil {
		util.Logger.LogError(err.Error())
		return
	}
	cnf, err := os.ReadFile(filepath.Join(path, "config", "config.json"))
	if err != nil {
		util.Logger.LogError(err.Error())
		return
	}
	err = json.Unmarshal(cnf, &config)
	if err != nil {
		util.Logger.LogError(err.Error())
		return
	}

	dbConf := mysql.DBConf{
		User:     config.DbUser,
		Password: config.DbPwd,
		URL:      config.DbHost,
		Schema:   config.DbName,
	}
	redisConf := redis.RedisConf{
		Address:  config.RedisAddr,
		Password: config.RedisPwd,
	}
	repos := []repository.InitRepo{
		&dbConf,
		&redisConf,
	}
	repository, err := repository.NewLoverRepository(repos)
	if err != nil {
		util.Logger.LogError(err.Error())
		return
	}

	service := service.NewLoverService(*repository)
	endpoint := endpoint.MakeLoverEndpoints(service)
	restServer := rest.NewRESTServer(endpoint)
	http.HandleFunc("/signup", restServer.SignUpHandler)
	http.HandleFunc("/login", restServer.LoginHandler)

	if err := http.ListenAndServe(":"+config.ServicePort, nil); err != nil {
		util.Logger.LogError(err.Error())
		return
	}
}
