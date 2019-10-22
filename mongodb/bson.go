package mongodb

import "go.mongodb.org/mongo-driver/bson"

func Bson2Map(v interface{}) map[string]interface{} {
	b, err := bson.Marshal(v)
	if err != nil {
		return nil
	}
	var m map[string]interface{}

	if err := bson.Unmarshal(b, &m); err != nil {
		return nil
	}

	delete(m, "_id")

	return m
}
