package dao

import (
	"github.com/globalsign/mgo"
	"github.com/ndphu/gm-api-golang/model"
	"github.com/globalsign/mgo/bson"
)

func ActorsCollection() *mgo.Collection {
	return Collection("actors")
}

func FindActorById(id string) (model.Actor, error) {
	actor := model.Actor{}
	err:=ActorsCollection().FindId(bson.ObjectIdHex(id)).One(&actor)
	return actor, err
}

func FindActorByKey(key string) (model.Actor, error) {
	actor := model.Actor{}
	err:=ActorsCollection().Find(bson.M{"key": key}).One(&actor)
	return actor, err
}
