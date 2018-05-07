package dao

import (
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/gm-api-golang/model"
)

func FindGenreById(id string) (*model.Genre, error) {
	var genre model.Genre
	err := Collection("genres").FindId(bson.ObjectIdHex(id)).One(&genre)
	return &genre, err
}
