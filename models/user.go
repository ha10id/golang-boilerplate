package models

import (
	"log"

	"../db"
)

//
type User struct {
	_id        string
	username   string
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
	log.Printf("Users %s", result)
	return result, nil
}
