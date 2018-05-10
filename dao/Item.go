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

func FindLatestMovies() ([]model.ItemLite, error) {
	var items []model.ItemLite
	err := ItemsCollection().Find(bson.M{"type": "MOVIE"}).Sort("-createdAt").Select(itemLiteSelector()).Limit(64).All(&items)
	return items, err
}

func FindLatestSeries() ([]model.ItemLite, error) {
	var items []model.ItemLite
	err := ItemsCollection().Find(bson.M{"type": "SERIE"}).Sort("-createdAt").Select(itemLiteSelector()).Limit(64).All(&items)
	return items, err
}

func ItemsCollection() *mgo.Collection {
	return Collection("items")
}

func FindItemsByGenre(genreTitle string, page int, size int) (bson.M, error) {
	var items []model.ItemLite
	query := ItemsCollection().Find(bson.M{
		"genres": genreTitle,
	}).Sort("-createdAt").Select(itemLiteSelector())

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
func itemLiteSelector() bson.M {
	return bson.M{
		"_id":      1,
		"title":    1,
		"subTitle": 1,
		"genres":   1,
		"actors":   1,
		"poster":   1,
	}
}

func FindItemsByActor(actorTitle string, page int, size int) (bson.M, error) {
	var items []model.ItemLite
	query := ItemsCollection().Find(bson.M{
		"actors": actorTitle,
	}).Sort("-createdAt").Select(itemLiteSelector())

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
