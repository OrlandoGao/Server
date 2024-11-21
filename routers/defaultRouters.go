package routers

import (
	controllers "ginproject/controllers"
	"github.com/gin-gonic/gin"
)

//默认页面路由初识化
func DefaultRountersInit(r *gin.Engine) {
	r.GET("/", controllers.MenuIndex)
}