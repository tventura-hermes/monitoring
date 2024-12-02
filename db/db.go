package db

import (
	"context"
	marketplace_domain "demo/api/marketplace/domain"
	mongo_db "demo/db/mongo"
	"errors"
	"fmt"
	"os"
)

type Database interface {
	Insert(result interface{}, collection string) error
	Select(collection string) ([]marketplace_domain.Message, error)
}

func NewDatabase(ctx context.Context) (Database, error) {
	dbType := os.Getenv("DB_CONTEXT")
	switch dbType {
	case "mongo":
		fmt.Print("Using mongo")
		return mongo_db.SetupMongo(ctx), nil
	default:
		return nil, errors.New("unsupported database backend")
	}
}
