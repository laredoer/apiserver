package router

import (
	"github.com/gin-gonic/gin"
	"apiserver/router/middleware"
	"net/http"
	"apiserver/handler/sd"
)

func Load(g *gin.Engine,mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Seruce)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound,"The incorrect API route")
	})

	u := g.Group("/v1/user")
	{
		u.POST("/:username", func(context *gin.Context) {

		})
	}

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk",sd.DiskCheck)
		svcd.GET("/cpu",sd.CPUCheck)
		svcd.GET("/ram",sd.RAMCheck)
	}

	return g
}