package db

import (
	"gopkg.in/mgo.v2"
)

var (
	mgoSession *mgo.Session
)

//
func Init() *mgo.Database {
	var err error

	mgoSession, err = mgo.Dial("mongodb://10.127.130.213:27017")

	if err != nil {
		panic(err)
	}

	return mgoSession.DB("appeals")

}

//
func CloseSession() {
	mgoSession.Close()
}
