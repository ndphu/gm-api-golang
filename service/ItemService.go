package service

import (
	"fmt"
	"errors"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
	"github.com/ndphu/gm-api-golang/model"
	"github.com/ndphu/gm-api-golang/dao"
)

func ReloadItem(item *model.Item) (*model.Item, error) {
	fmt.Println("reload item with type " + item.Type)
	if item.Type == "MOVIE" {
		return ReloadMovie(item)
	} else if item.Type == "SERIE" {
		return ReloadSerie(item)
	} else {
		return nil, errors.New("item type " + item.Type + " is not supported")
	}
}
func ReloadSerie(item *model.Item) (*model.Item, error) {
	return nil, nil
}
func ReloadMovie(item *model.Item) (*model.Item, error) {
	fmt.Println("removing all episodes for item " + item.Id.Hex())
	err := dao.RemoveEpisodesByItemId(item.Id.Hex())
	if err != nil && err != mgo.ErrNotFound {
		fmt.Printf("fail to remove existing episode for movie %s error = %v\n", item.Id.Hex(), err)
		return nil, err
	}
	fmt.Println("crawling from play url " + item.PlayUrl)
	videoSource, srt, err := CrawVideoSourceMQTT(item.PlayUrl)
	episode := model.Episode{
		Id: bson.NewObjectId(),
		Title: item.Title,
		SubTitle: item.SubTitle,
		CrawUrl: item.PlayUrl,
		ItemId: item.Id.Hex(),
		Order: 0,
		Srt: srt,
		VideoSource: videoSource,
	}
	err = dao.SaveEpisode(&episode)
	return item, err
}
