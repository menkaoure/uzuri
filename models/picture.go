package models

import "gopkg.in/mgo.v2/bson"

type picture struct {
	Id				bson.ObjectId `json:"id" bson:"_id"`
	Name			string		 `json:"name" bson:"name"`
}