package models

import (
	"context"
	"log"

	"../db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Department struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title    string             `json:"title" bson:"title"`
	SubTitle string             `json:"subtitle" bson:"subtitle"`
}

// GetAll user
func (h Department) GetAll() ([]*Department, error) {
	var results []*Department

	cur, err := db.Departments.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var elem Department
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
