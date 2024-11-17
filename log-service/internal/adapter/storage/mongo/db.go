package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/bedirhangull/hrcubo/log-service/internal/adapter/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client *mongo.Client
}

func New(config *config.DB, ctx context.Context) (*DB, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		config.Connection,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	log.Println("MongoDB URL: ", url)

	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: client,
	}, nil
}

func (d *DB) Close(ctx context.Context) error {
	return d.Client.Disconnect(ctx)
}
