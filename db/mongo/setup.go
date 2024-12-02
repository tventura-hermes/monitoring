package mongo_db

import (
	"context"
	errors_helpers "demo/helpers/errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client   *mongo.Client
	ctx      context.Context
	database *mongo.Database
}

func SetupMongo(ctx context.Context) *MongoClient {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI))

	if err != nil {
		errors_helpers.ReportError(ctx, err)
		panic("app stopped")
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		errors_helpers.ReportError(ctx, err)
		panic("db conexion failed")
	}

	name := os.Getenv("DB_NAME")
	database := client.Database(name)

	return &MongoClient{
		client:   client,
		ctx:      ctx,
		database: database,
	}
}
