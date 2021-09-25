package router

import (
	"github.com/gin-gonic/gin"

	"blog-api/blog-web/api"
)

func InitBlogRouter(Router *gin.RouterGroup) {
	BlogRouter := Router.Group("blog")

	{
		BlogRouter.GET("article", api.GetBlogList)
		BlogRouter.POST("article_create", api.ArticleCreate)
		BlogRouter.GET("article_detail", api.GetArticleDetail)
	}
}
