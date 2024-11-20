package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MenuIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "default/index.html", gin.H{})
}