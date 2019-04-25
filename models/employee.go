package models

import (
	"context"
	"log"

	"../db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Filter struct {
	Department primitive.ObjectID `json:"department" bson:"department,omitempty"`
}

type Employee struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Role  string             `json:"role" bson:"role"`
	Phone string             `json:"phone" bson:"phone"`
	Room  string             `json:"room" bson:"room"`
}

// GetAll user
func (h Employee) GetAll() ([]*Employee, error) {
	var results []*Employee

	cur, err := db.Employees.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var elem Employee
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

// Get user by ID
func (h Employee) Get(id string) (*Employee, error) {
	// filter := bson.D{primitive.E{Key: "_id", Value: id}}
	idB, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var result Employee
	err = db.Employees.FindOne(context.TODO(), bson.M{"_id": idB}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get employees by department
func (h Employee) GetByDepartment(depId string) ([]*Employee, error) {
	var results []*Employee
	depID, _ := primitive.ObjectIDFromHex(depId)
	filter := Filter{Department: depID}

	//= {"department": depId}
	log.Printf("DEPARTMENT: %s\nFILTER: %v", depId, filter)
	cur, err := db.Employees.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var elem Employee
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
