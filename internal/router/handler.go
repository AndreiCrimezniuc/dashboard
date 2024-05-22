package router

import (
	"dashboard/internal/dashboard"
	"dashboard/pkg/gen/openapi/dashboardapi"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	dashboardService dashboard.DashboardService
}

func New(dashboardAPI dashboard.DashboardService) RequestHandler {
	return RequestHandler{
		dashboardAPI,
	}
}

func (r RequestHandler) GetTrafficEarningEpoch(ctx *gin.Context) {
	_, err := r.dashboardService.GetV1Earnings(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	response := dashboardapi.GetTrafficEarningEpochResponse{
		JSON200: nil, // TODO fill from dashboardService
	}
	ctx.JSON(http.StatusOK, &response)
}

func (r RequestHandler) GetTrafficEarningLast30days(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
