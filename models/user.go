package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id				bson.ObjectId `json:"id" bson:"_id"`
	Email			string		 `json:"email" bson:"email"`
	PasswordHash	string		 `json:"password_hash" bson:"password_hash"`
}