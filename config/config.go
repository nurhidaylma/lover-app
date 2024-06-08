package config

type Service struct {
	ServicePort string `json:"serviceport"`
}

type Database struct {
	DbHost string `json:"dbhost"`
	DbName string `json:"dbname"`
	DbUser string `json:"dbuser"`
	DbPwd  string `json:"dbpwd"`
}

type Redis struct {
	RedisAddr string `json:"address"`
	RedisPwd  string `json:"password"`
}

type ServiceConfig struct {
	Service  `json:"service"`
	Database `json:"database"`
	Redis    `json:"redis"`
}
