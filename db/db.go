package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Package contains all db interactions

var (
	NasaDatabase   *mongo.Database
	ApodCollection *mongo.Collection
	Ctx            context.Context
	CtxCancel      context.CancelFunc
	Client         mongo.Client
	err            error
	dbUrl          string = "mongodb+srv://lordmalcher:Bartek222@cluster0-d545f.mongodb.net/test?retryWrites=true&w=majority"
)

func init() {
	Client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}
	Ctx, CtxCancel = context.WithCancel(context.Background())

	if err = Client.Connect(Ctx); err != nil {
		log.Fatal(err)
	}
	NasaDatabase = Client.Database("nasa")
	ApodCollection = NasaDatabase.Collection("apod")
}
