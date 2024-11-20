package routers

import (
	controllers "ginproject/controllers"

	"github.com/gin-gonic/gin"
)

//管理员路由初识化
func AdminRoutersInit(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		admin.GET("/", controllers.AdminIndex)
	}
} 