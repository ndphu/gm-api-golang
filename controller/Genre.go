package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/ndphu/gm-api-golang/model"
	"github.com/ndphu/gm-api-golang/dao"
)

func GenresController(g *gin.RouterGroup) {
	g.GET("", func(c *gin.Context) {
		var genres []model.Genre
		err := dao.Collection("genres").Find(nil).All(&genres)
		if err == nil {
			c.JSON(200, gin.H{
				"docs": genres,
				"total": len(genres),
			})
		} else {
			c.JSON(500, gin.H{"err": err})
		}
	})

	g.GET("/:id/items", func(c *gin.Context) {
		id := c.Param("id")
		genre, err := dao.FindGenreById(id)
		if err != nil {
			fmt.Println("could not find genre with id " + id)
			c.JSON(500, gin.H{"err": err})
			return
		}
		pg := Pagination{}
		if err = parsePagination(c, &pg); err != nil {
			fmt.Println("missing pagination query")
			c.JSON(500, gin.H{"err": err})
			return
		}

		items, err := dao.FindItemsByGenre(genre.Title, pg.Page, pg.Size)
		if err != nil {
			c.JSON(500, gin.H{"err": err})
		} else {
			c.JSON(200, gin.H{
				"genre": genre,
				"items": items,
			})
		}
	})
}
