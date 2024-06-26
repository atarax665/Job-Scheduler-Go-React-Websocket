package status

import (
	"scheduler-service/internal/service/status"

	"github.com/gin-gonic/gin"
)

type StatusController interface {
	HandleNewConnection(c *gin.Context)
}

type statusController struct {
	statusService status.StatusService
}
