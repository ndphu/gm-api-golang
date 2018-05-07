package controller

import (
	"github.com/gin-gonic/gin"
	"gm-api-golang/dao"
	"fmt"
	"gm-api-golang/service"
)

func ItemController(g *gin.RouterGroup) {
	g.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("find item by id " + id)
		item, err := dao.FindItemById(id)
		if err != nil {
			c.JSON(500, gin.H{"err": err})
		} else {
			c.JSON(200, item)
		}
	})

	g.GET("/:id/episodes", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("find episodes by item id " + id)
		episodes, err := dao.FindEpisodesByItemId(id)
		if err != nil {
			c.JSON(500, gin.H{"err": err})
		} else {
			c.JSON(200, episodes)
		}
	})

	g.GET("/:id/reload", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("reload item id " + id)
		item, err := dao.FindItemById(id)
		if err != nil {
			c.JSON(500, gin.H{"err": err})
		} else {
			fmt.Println(item.Title)
			item, err = service.ReloadItem(item)
			if err != nil {
				c.JSON(500, gin.H{"err": err})
			} else {
				c.JSON(200, item)
			}
		}
	})
}
