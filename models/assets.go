package models

import (
"context"
"log"

"../db"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
)

type Assets struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Assets      []interface{}           `json:"assets" bson:"assets"`
}

// GetAll assets
func (h Assets) GetAll(filter string) ([]*Assets, error) {
	var results []*Assets
	filterDB := bson.M{}
	if filter != "" {
		filterDB = bson.M{"name": filter}
	}
	cur, err := db.Assets.Find(context.TODO(), filterDB)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var elem Assets
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return results, nil
}
