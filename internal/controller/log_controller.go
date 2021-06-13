package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"securities_admin/internal/service"
	"strconv"
)

type ILogController interface {
	GetLogs(ctx *gin.Context)
}

type LogController struct {
	logService service.IRequestService
}

func NewLogController(logService service.IRequestService) *LogController {
	return &LogController{logService: logService}
}

func (c *LogController) GetLogs(ctx *gin.Context) {
	logType, err := strconv.Atoi(ctx.Param(RequestTypeParam))
	if err != nil {
		log.Printf("Failed to convert logType %s", err)
	}

	userID := ctx.Query(UserIDQuery)

	page, err := strconv.Atoi(ctx.Query(PageQuery))
	if err != nil {
		page = 1
		log.Printf("Failed to convert logType %s", err)
	}

	logs, err := c.logService.GetRequests(logType, page, userID)
	if err != nil {
		ctx.JSON(HTTPCodeServerError, gin.H{
			"message": "Get Logs failed",
		})
	}

	ctx.JSON(HTTPCodeOK, gin.H{
		"data": logs,
	})
}
