package service

import (
	"github.com/ndphu/gm-api-golang/dao"
	"github.com/ndphu/gm-api-golang/model"
	"github.com/globalsign/mgo/bson"
)

func SearchMovies(searchString string, page int, size int) (map[string]interface{}, error) {
	return searchItems(bson.M{
		"$or":  conditions(searchString),
		"type": "MOVIE",
	}, page, size)
}

func SearchSeries(searchString string, page int, size int) (map[string]interface{}, error) {
	return searchItems(bson.M{
		"$or":  conditions(searchString),
		"type": "SERIE",
	}, page, size)
}

func SearchActors(searchString string, page int, size int) (map[string]interface{}, error) {
	query  := dao.ActorsCollection().Find(bson.RegEx{
		Pattern: ".*" + searchString + ".*",
		Options: "i",
	})
	total, err := query.Count()
	if err != nil {
		return nil, err
	}
	var items []model.Item
	err = query.Skip((page - 1) * size).Limit(size).All(&items)
	if err != nil {
		return nil, err
	}
	return buildResult(items, total, size, page), err
}

func searchItems(search bson.M, page int, size int) (map[string]interface{}, error) {
	query := dao.ItemsCollection().Find(search)
	total, err := query.Count()
	if err != nil {
		return nil, err
	}
	var items []model.Item
	err = query.Skip((page - 1) * size).Limit(size).All(&items)
	if err != nil {
		return nil, err
	}
	return buildResult(items, total, size, page), err
}

func buildResult(items []model.Item, total int, size int, page int) map[string]interface{} {
	if items == nil {
		items = make([]model.Item, 0)
	}
	result := make(map[string]interface{})
	result["docs"] = items
	result["total"] = total
	result["limit"] = size
	result["page"] = page
	result["pages"] = total/size + 1
	return result
}
func conditions(searchString string) []bson.M {
	findItemCondition := []bson.M{
		{"title": bson.RegEx{Pattern: ".*" + searchString + ".*", Options: "i"}},
		{"normTitle": bson.RegEx{Pattern: ".*" + searchString + ".*", Options: "i"}},
		{"subTitle": bson.RegEx{Pattern: ".*" + searchString + ".*", Options: "i"}},
		{"normSubTitle": bson.RegEx{Pattern: ".*" + searchString + ".*", Options: "i"}},
	}
	return findItemCondition
}
