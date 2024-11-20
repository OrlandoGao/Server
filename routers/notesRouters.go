package routers

import (
	controllers "ginproject/controllers"

	"github.com/gin-gonic/gin"
)

//论坛路由初始化
func NotesRoutersInit(r *gin.Engine) {
	r.GET("/notes", controllers.NotesIndex)
}