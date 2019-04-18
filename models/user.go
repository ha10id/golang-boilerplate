package models

import (
	"../db"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//
type User struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserName   string        `json:"username" bson:"username"`
	firstName  string
	middleName string
	lastName   string
}

//
func (h User) GetAll() ([]User, error) {

	con := db.Init().C("users")

	defer db.CloseSession()

	var result []User

	err := con.Find(nil).All(&result)

	if err != nil {
		return nil, err
	}
	log.Printf("Users %v", result)
	return result, nil
}
