package mongo_db

import (
	marketplace_domain "demo/api/marketplace/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoClient) Select(collection string) ([]marketplace_domain.Message, error) {
	var results []marketplace_domain.Message

	cur, err := m.database.Collection(collection).Find(m.ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(m.ctx)

	for cur.Next(m.ctx) {
		var result marketplace_domain.Message
		err := cur.Decode(&result)

		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
