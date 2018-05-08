package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/gm-api-golang/dao"
	"github.com/ndphu/gm-api-golang/service"
)

func EpisodeController(g *gin.RouterGroup) {
	g.GET("/:id/reload", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("reload episode " + id)
		e, err := dao.FindEpisodeById(id)
		if err != nil {
			fmt.Printf("episode reloaded failed error = %v\n", err)
			c.JSON(500, gin.H{"err": err})
			return
		}
		videoSource, srt, err := service.CrawVideoSource(e.CrawUrl)
		if err != nil {
			fmt.Printf("episode reloaded failed error = %v\n", err)
			c.JSON(500, gin.H{"err": err})
			return
		}
		e.VideoSource = videoSource
		e.Srt = srt
		if err := dao.UpdateEpisode(e); err != nil {
			fmt.Printf("episode reloaded failed error = %v\n", err)
			c.JSON(500, gin.H{"err": err})
			return
		} else {
			fmt.Println("episode reloaded successfully")
			c.JSON(200, e)
		}
	})
}
