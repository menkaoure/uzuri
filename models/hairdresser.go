package models

import "gopkg.in/mgo.v2/bson"

const (
	Amateur = iota
	Pro
)

const (
	Haidress = iota
	Makeup
	HaidressAndMakeup
)

type Hairdresser struct {
		Id				bson.ObjectId 	`json:"id" bson:"_id"`
		Name			string		 	`json:"name" bson:"name"`
		User			string			`json:"user" bson:"user"`
		Status			int 			`json:"status" bson:"status"`
		Activity		int 			`json:"activity" bson:"activity"`
		Address			Address 		`json:"address" bson:"address"`
}