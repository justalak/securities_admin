package route

import (
	"github.com/gin-gonic/gin"
	"securities_admin/internal/controller"
)

type AdminServer struct {
	LogController controller.ILogController
	Router *gin.Engine
}

func NewAdminServer(logController controller.ILogController)*AdminServer {
	router := gin.Default()

	router.GET(GetLogs, logController.GetLogs)

	return &AdminServer{
		Router: router,
		LogController: logController,
	}
}


