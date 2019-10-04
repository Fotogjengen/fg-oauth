package handlers

import (
	"hilfling-oauth/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SecurityLevel struct {
	Id    int    `json:"id" binding:"required"`
	Level string `json:"level" binding:"required"`
}

func getSecurityLevels() ([]SecurityLevel, error) {
	const q = `SELECT * FROM security_level;`
	rows := database.Query(q)
	results := make([]SecurityLevel, 0)
	for rows.Next() { // Loop through all DB rows
		var id int
		var level string
		err := rows.Scan(&id, &level)
		if err != nil {
			return nil, err
		}
		results = append(results, SecurityLevel{id, level})
	}
	return results, nil
}

func GetSecurityLevels(ctx *gin.Context) {
	results, err := getSecurityLevels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "internal error: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, results)
}
