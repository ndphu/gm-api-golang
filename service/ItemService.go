package service

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/gm-api-golang/dao"
	"github.com/ndphu/gm-api-golang/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/ndphu/gm-api-golang/config"
	"github.com/ndphu/gm-api-golang/utils"
	"net/url"
	"net/http"
	"bytes"
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
	if item.Type != "SERIE" {
		return item, errors.New("unsupported item type " + item.Type)
	}
	fmt.Printf("reloading SERIE [%s] %s\n", item.Id.Hex(), item.Title)
	fmt.Println("serie url", item.PlayUrl)
	var doc *goquery.Document
	var err error
	if config.Get().UseCookie == true {
		playUrl, err :=url.Parse(item.PlayUrl)
		if err != nil {
			return item, err
		}
		cookie := http.Cookie{
			Name: "___cc",
			Value: "VNM",
		}
		rawData, err := utils.GetWithCookie(playUrl, &cookie, make(map[string]string))
		if err != nil {
			return item, err
		}
		doc, err = goquery.NewDocumentFromReader(bytes.NewReader(rawData))
	} else {
		doc, err = goquery.NewDocument(item.PlayUrl)
	}

	if err != nil {
		fmt.Println("fail to load goquery " + err.Error())
		return item, err
	}

	episodes := make([]*model.Episode,0)

	doc.Find(".episode-main li a").Each(func(idx int, s *goquery.Selection) {
		episodeUrl := s.AttrOr("href", "")
		fmt.Println("found episode " + episodeUrl)
		episodes = append(episodes, &model.Episode{
			Id:       bson.NewObjectId(),
			Title:    fmt.Sprintf("Tập %d", idx + 1),
			SubTitle: fmt.Sprintf("Episode %d", idx + 1),
			CrawUrl:  episodeUrl,
			Order:    idx,
			ItemId:   item.Id.Hex(),
		})
	})

	if len(episodes) > 0 {
		err := dao.RemoveEpisodesByItemId(item.Id.Hex())
		if err != nil {
			fmt.Println("fail to remove associated episodes")
			return item, err
		}
		err = dao.SaveAllEpisodes(episodes)
		if err != nil {
			fmt.Println("fail to save associated episodes", err.Error())
			return item, err
		}
	}

	return item, err
}

func ReloadMovie(item *model.Item) (*model.Item, error) {
	fmt.Println("removing all episodes for item " + item.Id.Hex())
	err := dao.RemoveEpisodesByItemId(item.Id.Hex())
	if err != nil && err != mgo.ErrNotFound {
		fmt.Printf("fail to remove existing episode for movie %s error = %v\n", item.Id.Hex(), err)
		return nil, err
	}
	fmt.Println("crawling from play url " + item.PlayUrl)
	videoSource, srt, err := CrawVideoSource(item.PlayUrl)
	episode := model.Episode{
		Id:          bson.NewObjectId(),
		Title:       item.Title,
		SubTitle:    item.SubTitle,
		CrawUrl:     item.PlayUrl,
		ItemId:      item.Id.Hex(),
		Order:       0,
		Srt:         srt,
		VideoSource: videoSource,
	}
	err = dao.SaveEpisode(&episode)
	return item, err
}
