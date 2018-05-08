package model

import "github.com/globalsign/mgo/bson"

type Actor struct {
	Id    bson.ObjectId `json:"_id" bson:"_id"`
	Title string        `json:"title" bson:"title"`
	Key   string        `json:"key" bson:"key"`
}
