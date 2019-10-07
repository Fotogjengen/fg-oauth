package handlers

import (
	"hilfling-oauth/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type position struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func getPositions() ([]position, error) {
	const q = `SELECT * FROM position;`
	rows := database.Query(q)
	results := make([]position, 0)
	for rows.Next() { // Loop through all DB rows
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		results = append(results, position{id, name})
	}
	return results, nil
}

func GetPositions(ctx *gin.Context) {
	results, err := getPositions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "internal error: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, results)
}
