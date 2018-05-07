package model

import "github.com/globalsign/mgo/bson"

type Episode struct {
	Id          bson.ObjectId `json:"_id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	SubTitle    string        `json:"subTitle" bson:"subTitle"`
	Order       int           `json:"order" bson:"order"`
	CrawUrl     string        `json:"crawUrl" bson:"crawUrl"`
	ItemId      string        `json:"itemId" bson:"itemId"`
	Srt         string        `json:"srt" bson:"srt"`
	VideoSource string        `json:"videoSource" bson:"videoSource"`
}
