package db

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "users"
)

func GetDBCollection() (*mongo.Collection, error) {
	log.SetFormatter(&log.JSONFormatter{})

	// client, err := mongo.NewClient((options.Client().ApplyURI("mongodb://localhost:27017")))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://standard:<password>@cluster0.pdpui.mongodb.net/<database>?retryWrites=true&w=majority",
	))

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	//Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	collection := client.Database("GoLogin").Collection("users")
	log.Info("Connection Success!!!")
	return collection, nil
}
