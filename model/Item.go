package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Item struct {
	Id           bson.ObjectId `json:"_id" bson:"_id"`
	Directors    []string      `json:"directors" bson:"directors"`
	Actors       []string      `json:"actors" bson:"actors"`
	Genres       []string      `json:"genres" bson:"genres"`
	Countries    []string      `json:"countries" bson:"countries"`
	Title        string        `json:"title" bson:"title"`
	NormTitle    string        `json:"normTitle" bson:"normTitle"`
	SubTitle     string        `json:"subTitle" bson:"subTitle"`
	NormSubTitle string        `json:"normSubTitle" bson:"normSubTitle"`
	Poster       string        `json:"poster" bson:"poster"`
	BigPoster    string        `json:"bigPoster" bson:"bigPoster"`
	Source       string        `json:"source" bson:"source"`
	Hash         string        `json:"hash" bson:"hash"`
	PlayUrl      string        `json:"playUrl" bson:"playUrl"`
	Type         string        `json:"type" bson:"type"`
	ReleaseData  time.Time     `json:"releaseDate" bson:"releaseDate"`
	Content      string        `json:"content" bson:"content"`
	CreatedAt    time.Time     `json:"createdAt" bson:"createdAt"`
}

func (i *Item) Save() (*Item) {
	return i
}
