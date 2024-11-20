package routers

import (
	controllers "ginproject/controllers"

	"github.com/gin-gonic/gin"
)

//论坛路由初始化
func ForumRoutersInit(r *gin.Engine) {
	forum := r.Group("/forum")
	{
		forum.GET("/", controllers.ForumIndex)
		forum.GET("/login", controllers.ForumLogin)
	}
}