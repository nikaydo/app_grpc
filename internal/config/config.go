package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Env struct {
	EnvMap            map[string]string
	TimeoutConnect    int
	TimeoutDisconnect int
	TimeoutRead       int
	TimeoutWrite      int
}

func ReadEnv() Env {
	envMap, err := godotenv.Read("./.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	return Env{EnvMap: envMap}
}
