package service

import (
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/gm-api-golang/dao"
	"github.com/ndphu/gm-api-golang/model"
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
	query := dao.ActorsCollection().Find(bson.M{
		"title": bson.RegEx{
			Pattern: ".*" + searchString + ".*",
			Options: "i",
		}})
	total, err := query.Count()
	if err != nil {
		return nil, err
	}
	var actors []model.Actor
	err = query.Skip((page - 1) * size).Limit(size).All(&actors)
	if err != nil {
		return nil, err
	}
	if actors == nil {
		return emptyResult(), err
	} else {
		return buildResult(actors, total, size, page), err
	}
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
	if items == nil {
		return emptyResult(), err
	} else {
		return buildResult(items, total, size, page), err
	}
}

func emptyResult() map[string]interface{} {
	result := make(map[string]interface{})
	result["docs"] = make([]model.Item, 0)
	result["total"] = 0
	result["limit"] = 0
	result["page"] = 0
	result["pages"] = 0
	return result
}

func buildResult(items interface{}, total int, size int, page int) map[string]interface{} {
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
