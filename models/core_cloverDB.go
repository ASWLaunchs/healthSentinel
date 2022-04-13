package models

import (
	clover "github.com/ostafen/clover"
)

var DBCloverDB *clover.DB

func init() {
	db, err := clover.Open("data/sql/healthSentinel")
	if err != nil {
		panic(err.Error())
	}
	DBCloverDB = db
}

//CreateDB() if not exist db , it will be created.
func CreateDB() {
	exist, err := DBCloverDB.HasCollection("studentList")
	if err != nil {
		panic(err.Error())
	}

	if !exist {
		err = DBCloverDB.CreateCollection("studentList")
		if err != nil {
			panic(err.Error())
		}
	}
}
