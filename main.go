package main

import (
	"github.com/0xfzz/tuwitt/bootstrap"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	bootstrap.Start()
}
