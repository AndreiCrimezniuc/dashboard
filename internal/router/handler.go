package router

import (
	"dashboard/internal/dashboard"
	"dashboard/pkg/gen/openapi/dashboardapi"
	openapi_types "github.com/SealNTibbers/oapi-codegen/pkg/types"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	dashboardService dashboard.DashboardService
	logger           *slog.Logger // toDo: add logger interface instead
}

func New(dashboardAPI dashboard.DashboardService) RequestHandler {
	return RequestHandler{
		dashboardService: dashboardAPI,
		logger:           slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func (r RequestHandler) GetTrafficEarningEpoch(c *gin.Context) {
	c.Set("user_id", 1) // jwt mocking

	earningAmount, err := r.dashboardService.GetV1Earnings(c)
	if err != nil {
		r.logger.Error("error in getting earnings epoch:" + err.Error())
		c.JSON(http.StatusBadRequest, nil)
	}

	response := dashboardapi.GetTrafficEarningEpochResponse{
		JSON200: &struct {
			Earnings *float32 `json:"earnings,omitempty"`
		}{
			Earnings: &earningAmount,
		},
	}
	c.JSON(http.StatusOK, &response)
}

func (r RequestHandler) GetTrafficEarningLast30days(c *gin.Context) {
	c.Set("user_id", 1) // jwt mocking

	earningAmount, err := r.dashboardService.GetV1EarningsStats(c)
	if err != nil {
		r.logger.Error("error in getting earnings epoch:" + err.Error())
		c.JSON(http.StatusBadRequest, nil)
	}

	earningAmountOpenApiFormat := make([]struct { // that is weird. Can we just change JSON200 struct?
		Date     *openapi_types.Date `json:"date,omitempty"`
		Earnings *float32            `json:"earnings,omitempty"`
	}, 0, len(earningAmount))

	for _, val := range earningAmount {

	}

	response := dashboardapi.GetTrafficEarningLast30daysResponse{
		JSON200: earningAmount,
	}

	c.JSON(http.StatusOK, &response)
}
