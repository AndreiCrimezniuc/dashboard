package dashboard

import "context"

type dashboardServiceImpl struct {
}

type DashboardService interface {
	GetV1Earnings(ctx context.Context) (float32, error)
	GetV1EarningsStats(ctx context.Context) (string, error)
}

func Init() (DashboardService, error) {
	return &dashboardServiceImpl{}, nil
}

func (d dashboardServiceImpl) GetV1Earnings(ctx context.Context) (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (d dashboardServiceImpl) GetV1EarningsStats(ctx context.Context) (string, error) {
	//TODO implement me
	panic("implement me")
}
