package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotesIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "notes/designPatterns.html", gin.H{})
}