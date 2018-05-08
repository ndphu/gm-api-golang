package dao

import (
	"github.com/globalsign/mgo"
)

func ActorsCollection() *mgo.Collection {
	return Collection("items")
}
