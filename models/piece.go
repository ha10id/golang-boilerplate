package models

import (
	"context"
	"log"

	"../db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Piece struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Body        string             `json:"body" bson:"body"`
	SubTitle    string             `json:"subtitle" bson:"subtitle"`
	Assets      []string           `json:"assets" bson:"assets"`
	CreatedAt   primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt   primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	PublishedAt primitive.DateTime `json:"publishedAt" bson:"publishedAt"`
}

// GetAll user
func (h Piece) GetAll() ([]*Piece, error) {
	var results []*Piece

	cur, err := db.News.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var elem Piece
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
