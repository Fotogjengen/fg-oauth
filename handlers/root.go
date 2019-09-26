package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type path struct {
	Name string
	Path string
}

var paths = []path{
	{
		Name: "Test",
		Path: "test",
	},
}

func GetRoot(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api-root.tmpl", gin.H{
		"paths": paths,
	})
}
