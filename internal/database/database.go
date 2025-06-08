package database

import (
	"log"
	"main/internal/config"
)

type Database struct {
	Video VideoDB

	ENV config.Env
}

func InitDB(Env config.Env) (Database, error) {
	video, err := VideoInit(Env)
	if err != nil {
		return Database{}, err
	}
	log.Println("MongoDB connected.")

	return Database{Video: video, ENV: Env}, nil
}
