package controller

import (
	"github.com/gin-gonic/gin"
	"gm-api-golang/dao"
	"gm-api-golang/model"
	"sync"
)

type Section struct {
	Category model.Genre  `json:"category"`
	Items    []model.Item `json:"items"`
}

func HomeController(g *gin.RouterGroup) {
	g.GET("", func(c *gin.Context) {
		wg := sync.WaitGroup{}
		var movieSection Section
		var serieSection Section
		var err1 error
		var err2 error
		wg.Add(1)
		go func() {
			movieSection, err1 = getLatestMovies()
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			serieSection, err2 = getLatestSeries()
			wg.Done()
		}()
		wg.Wait()
		if err1 != nil {
			c.JSON(500, gin.H{"err": err1})
		} else if err2 != nil {
			c.JSON(500, gin.H{"err": err2})
		} else {
			c.JSON(200, []Section{movieSection, serieSection})
		}
	})
}

func getLatestMovies() (Section, error) {
	var items []model.Item
	items, err:= dao.FindLatestMovies()
	return Section{
		Category: model.Genre{
			Key:   "latest-movies",
			Title: "Phim Lẻ Mới Cập Nhật",
		},
		Items: items,
	}, err
}

func getLatestSeries() (Section, error) {
	var items []model.Item
	items, err:= dao.FindLatestSeries()
	return Section{
		Category: model.Genre{
			Key:   "latest-series",
			Title: "Phim Bộ Mới Cập Nhật",
		},
		Items: items,
	}, err
}
