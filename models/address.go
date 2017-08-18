package models

import "gopkg.in/mgo.v2/bson"

type Address struct {
	Id				bson.ObjectId `json:"id" bson:"_id"`
	Street			string		 `json:"street" bson:"street"`
	StreetNumber	string		 `json:"street_number" bson:"street_number"`
	Zip				string		 `json:"zip" bson:"zip"`
	City			string		 `json:"city" bson:"city"`
	Country			string		 `json:"country" bson:"country"`
}
