package models

import "gopkg.in/mgo.v2/bson"

const (
	Male = iota
	Female
	Mix
)

type Ad struct {
	Id				bson.ObjectId `json:"id" bson:"_id"`
	Name			string		 `json:"name" bson:"name"`
	Description		string		 `json:"description" bson:"description"`
	Owner			string	 	 `json:"hairdresser" bson:"hairdresser"` // Hairdresser
	Pictures		string		 `json:"pictures" bson:"pictures"`
	CustomerGender	string		 `json:"customer_gender" bson:"customer_gender"`
}
