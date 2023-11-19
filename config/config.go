package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type dbconfig struct {
	Name     string
	Username string
	Host     string
	Password string
	Port     int
}
type appconfig struct {
	Port int
	Url  string
}
type jwtconifg struct {
	Secret string
}
type Config struct {
	DB  dbconfig
	APP appconfig
	JWT jwtconifg
}

func Get() Config {
	_ = godotenv.Load("../.env")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(fmt.Sprintf("Error while converting database port : %s", err))
	}
	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		panic(fmt.Sprintf("Error while converting app port : %s", err))
	}

	return Config{
		DB: dbconfig{
			Name:     os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USERNAME"),
			Host:     os.Getenv("DB_HOST"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     dbPort,
		},
		APP: appconfig{
			Port: appPort,
			Url:  os.Getenv("APP_URL"),
		},
		JWT: jwtconifg{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}
