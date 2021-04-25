// Cyb3rVector EscapePod Proxy
//  by cyb3rdog
//
// (c) 2021 Vaclav Macha
// All rights reserved.
//
// mongoclient - handles the communication with mongo db
//
package mongoclient

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	db "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	client  *mongo.Client
	intents *mongo.Collection
}

// New returns a populated MongoClient struct
func New(opts ...Option) (*MongoClient, error) {
	cfg := options{}

	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.host == "" {
		return nil, fmt.Errorf("no mongo db host defined")
	}
	if cfg.port == "" {
		cfg.port = "27017"
	}
	if cfg.user == "" {
		return nil, fmt.Errorf("no mongo db user defined")
	}
	if cfg.pass == "" {
		return nil, fmt.Errorf("no mongo db password defined")
	}
	if cfg.name == "" {
		return nil, fmt.Errorf("no mongo db database name defined")
	}

	credential := db.Credential{
		AuthSource:  "admin",
		Username:    cfg.user,
		Password:    cfg.pass,
		PasswordSet: true,
	}

	// Set client options
	db_url := fmt.Sprintf("mongodb://%s:%s", cfg.host, cfg.port)
	clientOpts := db.Client().ApplyURI(db_url).SetAuth(credential).SetDirect(true)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to MongoDB
	db_client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = db_client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	intents := db_client.Database(cfg.name).Collection("intents")
	client := MongoClient{
		client:  db_client,
		intents: intents,
	}
	log.Printf("Successfully connected to MongoDB: %v", db_url)

	return &client, nil
}
