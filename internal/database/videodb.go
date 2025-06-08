package database

import (
	"context"
	"main/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoDB struct {
	Client *mongo.Client
	ENV    config.Env
}

func VideoInit(env config.Env) (VideoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(env.TimeoutConnect)*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(env.EnvMap["MANGODB_URL"])
	var db VideoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return db, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return db, err
	}
	db.Client = client
	db.ENV = env
	return db, nil
}

func (db *VideoDB) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(db.ENV.TimeoutDisconnect)*time.Second)
	defer cancel()
	err := db.Client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *VideoDB) AddVideo(title, Collection string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(db.ENV.TimeoutWrite)*time.Second)
	defer cancel()
	collection := db.Client.Database(db.ENV.EnvMap["MANGODB_NAME"]).Collection(Collection)
	doc := bson.D{
		{Key: "title", Value: title},
	}
	_, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	return nil
}
