package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/gm-api-golang/service"
)

func SearchController(g *gin.RouterGroup) {
	g.GET("/q/:query", func(c *gin.Context) {
		query := c.Param("query")
		fmt.Println("perform search with " + query)
		pg := Pagination{}
		if err := parsePagination(c, &pg); err != nil {
			fmt.Println("missing pagination query")
			c.JSON(500, gin.H{"err": err})
			return
		}
		movies, err := service.SearchMovies(query, pg.Page, pg.Size)
		if err != nil {
			c.JSON(500, gin.H{"err": err})
			return
		}

		series, err := service.SearchSeries(query, pg.Page, pg.Size)
		if err != nil {
			c.JSON(500, gin.H{"err": err})
			return
		}

		actors, err := service.SearchActors(query, pg.Page, pg.Size)
		if err != nil {
			c.JSON(500, gin.H{"err": err})
			return
		}

		c.JSON(200, gin.H{
			"movie": movies,
			"serie": series,
			"actor": actors,
		})
	})
}
