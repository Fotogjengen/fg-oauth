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
		Path: "path",
	},
	{
		Name: "Test2",
		Path: "Path2",
	},
	{
		Name: "Test2",
		Path: "Test2",
	},
}

func GetRoot(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api-root.tmpl", gin.H{
		"paths": paths,
	})
}
