package initialize

import (
	"github.com/gin-gonic/gin"

	"blog-api/blog-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	ApiGroup := Router.Group("v1")
	router.InitBlogRouter(ApiGroup)

	return Router
}
