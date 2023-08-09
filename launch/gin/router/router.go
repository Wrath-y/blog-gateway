package router

import (
	"gateway/infrastructure/common/context"
	"gateway/infrastructure/common/errcode"
	"gateway/launch/gin/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Recovery)
	r.Use(middleware.SetV)
	r.Use(context.Handle(middleware.CORS))
	r.NoRoute(NoRoute)

	r.Any("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	g := r.Group("/")
	loadApi(g)

	return r
}

func NoRoute(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, errcode.LibNoRoute)
}
