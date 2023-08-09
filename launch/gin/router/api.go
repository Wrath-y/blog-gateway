package router

import (
	"gateway/infrastructure/common/context"
	"gateway/interfaces/facade"
	"gateway/launch/gin/middleware"
	"github.com/gin-gonic/gin"
)

func loadApi(r *gin.RouterGroup) {
	a := r.Group("/api", context.Handle(middleware.Logging))
	{
		ars := a.Group("articles")
		{
			ars.GET("", context.Handle(facade.GetArticles))
			ars.GET("/a", context.Handle(facade.GetAllArticles))
		}
		ar := a.Group("article")
		{
			ar.GET("/:id", context.Handle(facade.GetArticle))
		}
	}
}
