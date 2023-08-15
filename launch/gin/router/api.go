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
		a.GET("/pixivs", context.Handle(facade.GetPixivs))
		a.GET("/friends", context.Handle(facade.GetFriends))

		ars := a.Group("articles")
		{
			ars.GET("", context.Handle(facade.GetArticles))
			ars.GET("/a", context.Handle(facade.GetAllArticles))
		}
		ar := a.Group("article")
		{
			ar.GET("/:id", context.Handle(facade.GetArticle))
		}

		cm := a.Group("comments")
		{
			cm.GET("", context.Handle(facade.GetComments))
			cm.GET("/count", context.Handle(facade.GetCommentCount))
			cm.POST("", context.Handle(facade.AddComment))
		}
	}
}
