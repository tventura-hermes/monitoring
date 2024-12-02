package mongo_db

import "fmt"

func (m *MongoClient) Insert(result interface{}, collection string) error {
	insert, err := m.database.Collection(collection).InsertOne(m.ctx, result)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(insert)

	return nil
}
