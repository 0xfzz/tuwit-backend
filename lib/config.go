package lib

import (
	"fmt"
	"os"
	"strconv"
)

type DBConfig struct {
	Name     string
	Username string
	Host     string
	Password string
	Port     int
}
type APPConfig struct {
	Port int
}
type JWTConfig struct {
	Secret string
}
type Config struct {
	DB  DBConfig
	APP APPConfig
	JWT JWTConfig
}

func NewConfig() Config {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(fmt.Sprintf("Error while converting database port : %s", err))
	}
	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		panic(fmt.Sprintf("Error while converting app port : %s", err))
	}

	return Config{
		DB: DBConfig{
			Name:     os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USERNAME"),
			Host:     os.Getenv("DB_HOST"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     dbPort,
		},
		APP: APPConfig{
			Port: appPort,
		},
		JWT: JWTConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}

// DB_NAME=
// DB_USERNAME=
// DB_HOST=
// DB_PASSWORD=
// DB_PORT=

// APP_PORT=
// JWT_SECRET=
