package health

import (
	db "github.com/yury-herlovich/go-rest-api/src/common"

	"net/http"

	"github.com/gin-gonic/gin"
)

var version = "0.0.1"

type Health struct {
	Version  string `json:"version"`
	Status   string `json:"status"`
	DbStatus string `json:"dbStatus"`
}

func HealthCheck(ctx *gin.Context) {
	dbStatus := dbCheck()

	status := "ok"

	if dbStatus != "ok" {
		status = "error"
	}

	health := Health{
		Version:  version,
		Status:   status,
		DbStatus: dbStatus,
	}

	ctx.IndentedJSON(http.StatusOK, health)
}

func dbCheck() string {
	dbStatus := "ok"

	if dbError := db.Check(); dbError != nil {
		dbStatus = "error"
	}

	return dbStatus
}
