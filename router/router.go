package router

import (
	"apiserver/handler/sd"
	"apiserver/handler/user"
	"apiserver/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect APi route.")
	})

	g.POST("/login", user.Login)

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("/create", user.Create)
		u.POST("/delete/:id", user.Delete)
		u.POST("/update/:id", user.Update)
		u.GET("/list", user.List)
		u.GET("/get/:username", user.Get)
	}

	//The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		//svcd.GET("/disk", sd.DiskCheck)
		//svcd.GET("/cpu", sd.CPUCheck)
		//svcd.GET("/ram", sd.RamCheck)
	}
	return g
}
