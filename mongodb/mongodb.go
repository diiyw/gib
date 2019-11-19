package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var mongoClients = make(map[string]*mongo.Client, 0)

type MgoConfig struct {
	Uri string `yaml:"uri"`
	DB  string `yaml:"db"`
}

func NewMongo(uri string, passwordSet bool) (*mongo.Client, error) {

	if client, ok := mongoClients[uri]; ok {
		return client, nil
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI(uri)
	clientOptions.Auth.PasswordSet = passwordSet

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	mongoClients[uri] = client

	return client, nil
}
