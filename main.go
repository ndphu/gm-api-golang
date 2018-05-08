package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/gm-api-golang/config"
	"github.com/ndphu/gm-api-golang/controller"
)

func main() {
	if config.Get().GinDebug == true {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	setupCORS(r)
	initControllers(r.Group("/api"))
	r.Run()
}
func initControllers(g *gin.RouterGroup) {
	controller.HomeController(g.Group("/home"))
	controller.GenresController(g.Group("/genre"))
	controller.ItemController(g.Group("/item"))
	controller.EpisodeController(g.Group("/episode"))
	controller.SearchController(g.Group("/search"))
	controller.ActorController(g.Group("/actor"))
	fmt.Println("controllers initialized")
}
func setupCORS(r *gin.Engine) {
	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	c.AllowCredentials = true
	c.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	c.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Content-Length", "X-Requested-With"}
	r.Use(cors.New(c))
}
