package models

import (
	"../db"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

//
type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	OID         string             `json:"OID" bson:"OID"`
	OIDS        []string           `json:"OIDS" bson:"OIDS"`
	Permission  int                `json:"permission" bson:"permission"`
	Goverment   primitive.ObjectID `json:"goverment" bson:"goverment"`
	UserName    string             `json:"username" bson:"username"`
	FirstName   string             `json:"firstName" bson:"firstName"`
	MiddleName  string             `json:"middleName" bson:"middleName"`
	LastName    string             `json:"lastName" bson:"lastName"`
	CreatedDate primitive.DateTime `json:"createdDate" bson:"createdDate"`
	Avatar      string             `json:"avatar" bson:"avatar"`
	Email       string             `json:"email" bson:"email"`
	Hash        string             `json:"-" bson:"hash"`
}

//
// GetAll user
func (h User) GetAll() ([]*User, error) {
	var results []*User

	cur, err := db.Users.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var elem User
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
