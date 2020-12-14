package util

import (
	"gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

)

func GetPage(c *gin.Context) int {

	//从context中获取page参数
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	//pagesize从setting中获取
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}