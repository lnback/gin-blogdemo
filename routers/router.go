package routers

import (
	"gin-blog/pkg/setting"
	v1 "gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter()  * gin.Engine{
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tags",v1.GetTags)
		apiV1.POST("/tags",v1.AddTag)
		apiV1.DELETE("/tags/:id",v1.DeleteTag)
		apiV1.PUT("/tags/:id",v1.EditTag)


		apiV1.GET("/articles",v1.GetArticles)
		apiV1.GET("/articles/:id",v1.GetArticle)
		apiV1.POST("/articles",v1.AddArticle)
		apiV1.DELETE("/articles/:id",v1.DeleteArticles)
		apiV1.PUT("/articles/:id",v1.EditArticles)
	}

	return r
}