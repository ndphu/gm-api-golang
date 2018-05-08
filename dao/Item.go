package dao

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/gm-api-golang/model"
)

func FindItemById(id string) (*model.Item, error) {
	var item model.Item
	err := Collection("items").FindId(bson.ObjectIdHex(id)).One(&item)
	return &item, err
}

func FindLatestMovies() ([]model.Item, error) {
	var items []model.Item
	err := ItemsCollection().Find(bson.M{"type": "MOVIE"}).Sort("-createdAt").Limit(64).All(&items)
	return items, err
}

func FindLatestSeries() ([]model.Item, error) {
	var items []model.Item
	err := ItemsCollection().Find(bson.M{"type": "SERIE"}).Sort("-createdAt").Limit(64).All(&items)
	return items, err
}

func ItemsCollection() *mgo.Collection {
	return Collection("items")
}

func FindItemsByGenre(genreTitle string, page int, size int) (bson.M, error) {
	var items []model.Item
	query := ItemsCollection().Find(bson.M{
		"genres": genreTitle,
	}).Sort("-createdAt")

	total, err := query.Count()
	if err != nil {
		return nil, err
	}
	pages := total/size + 1

	query.Skip((page - 1) * size).Limit(size).All(&items)
	return bson.M{
		"docs":  items,
		"total": total,
		"limit": size,
		"page":  page,
		"pages": pages,
	}, nil
}

func FindItemsByActor(actorTitle string, page int, size int) (bson.M, error) {
	var items []model.Item
	query := ItemsCollection().Find(bson.M{
		"actors": actorTitle,
	}).Sort("-createdAt")

	total, err := query.Count()
	if err != nil {
		return nil, err
	}
	pages := total/size + 1

	query.Skip((page - 1) * size).Limit(size).All(&items)
	return bson.M{
		"docs":  items,
		"total": total,
		"limit": size,
		"page":  page,
		"pages": pages,
	}, nil
}
