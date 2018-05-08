package dao

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/gm-api-golang/model"
)

func FindEpisodeById(id string) (*model.Episode, error) {
	var item model.Episode
	err := EpisodesCollection().FindId(bson.ObjectIdHex(id)).One(&item)
	return &item, err
}
func FindEpisodesByItemId(itemId string) ([]model.Episode, error) {
	var episodes []model.Episode
	err := EpisodesCollection().Find(bson.M{"itemId": itemId}).All(&episodes)
	return episodes, err
}

func EpisodesCollection() *mgo.Collection {
	return Collection("episodes")
}

func RemoveEpisodesByItemId(itemId string) error {
	return EpisodesCollection().Remove(bson.M{"itemId": itemId})
}

func SaveEpisode(episode *model.Episode) error {
	return EpisodesCollection().Insert(episode)
}

func UpdateEpisode(e *model.Episode) error {
	return EpisodesCollection().UpdateId(e.Id, e)
}
