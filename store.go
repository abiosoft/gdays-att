package main

import (
	"labix.org/v2/mgo"
	_ "labix.org/v2/mgo/bson"
)

var Database *mgo.Database

func init() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	Database = session.DB("att")

	err = Database.C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})

	if err != nil {
		panic(err)
	}
}

type User struct {
	Email string
}

func markAttendance(email string) error {
	user := &User{
		Email: email,
	}

	return Database.C("users").Insert(user)
}

func getAttendeesCount() int {
	count, err := Database.C("users").Count()
	if err != nil {
		panic(err)
	}
	return count
}
