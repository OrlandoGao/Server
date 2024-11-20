package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForumIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "forum/index.html", gin.H{})
}

func ForumLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "forum/login.html", gin.H{})
}