package dao

import (
	"gm-api-golang/model"
	"github.com/globalsign/mgo/bson"
)

func FindGenreById(id string) (*model.Genre, error) {
	var genre model.Genre
	err := Collection("genres").FindId(bson.ObjectIdHex(id)).One(&genre)
	return &genre, err
}
