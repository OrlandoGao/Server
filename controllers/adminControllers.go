package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/index.html", gin.H{})
}