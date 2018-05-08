package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ndphu/gm-api-golang/dao"
	"fmt"
)

func ActorController(g *gin.RouterGroup) {
	g.GET("/byKey/:key", func(c *gin.Context) {
		key := c.Param("key")
		fmt.Println("find actor by key " + key)
		actor, err := dao.FindActorByKey(key)
		if err != nil {
			c.JSON(500, gin.H{"err": err})
		} else {
			c.JSON(200, actor)
		}
	})

	g.GET("/byKey/:key/items", func(c *gin.Context) {
		key := c.Param("key")
		actor, err := dao.FindActorByKey(key)
		if err != nil {
			fmt.Println("could not find actor with key " + key)
			c.JSON(500, gin.H{"err": err})
			return
		}
		pg := Pagination{}
		if err = parsePagination(c, &pg); err != nil {
			fmt.Println("missing pagination query")
			c.JSON(500, gin.H{"err": err})
			return
		}

		items, err := dao.FindItemsByActor(actor.Title, pg.Page, pg.Size)
		if err != nil {
			c.JSON(500, gin.H{"err": err})
		} else {
			c.JSON(200, gin.H{
				"actor": actor,
				"items": items,
			})
		}
	})
}
