package config

import (
	"log"
	"strconv"

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

func (db *Env) SetTimeout() error {
	t, err := strconv.Atoi(db.EnvMap["MANGODB_TIMEOUT_CONNECT"])
	if err != nil {
		return err
	}
	db.TimeoutConnect = t
	t, err = strconv.Atoi(db.EnvMap["MANGODB_TIMEOUT_READ"])
	if err != nil {
		return err
	}
	db.TimeoutRead = t
	t, err = strconv.Atoi(db.EnvMap["MANGODB_TIMEOUT_WRITE"])
	if err != nil {
		return err
	}
	db.TimeoutWrite = t
	t, err = strconv.Atoi(db.EnvMap["MANGODB_TIMEOUT_DISCONNECT"])
	if err != nil {
		return err
	}
	db.TimeoutDisconnect = t
	return nil
}
