package router

import (
	"dashboard/internal/dashboard"
	"dashboard/pkg/gen/openapi/dashboardapi"

	"github.com/gin-gonic/gin"
)

func SetupRouter(dashboard dashboard.DashboardService) *gin.Engine {
	router := gin.Default()

	dashboardapi.RegisterHandlersWithOptions(router, New(dashboard), dashboardapi.GinServerOptions{})
	return router
}
