package dashboard

import (
	"context"
	"dashboard/internal/config"
	"errors"
	"fmt"
	openapi_types "github.com/SealNTibbers/oapi-codegen/pkg/types"
	"github.com/jackc/pgx/v5"
)

const (
	_allowedCurrency   = "USD"
	_statsNumberOfDays = 15
)

type dashboardServiceImpl struct {
	connection *pgx.Conn
}

type EarningStats struct {
	Date     openapi_types.Date
	Earnings *float32
}

type DashboardService interface {
	GetV1Earnings(ctx context.Context) (float32, error)
	GetV1EarningsStats(ctx context.Context) ([]EarningStats, error)
}

func Init(config config.Database) (DashboardService, error) {
	connection, er := pgx.Connect(
		context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Password, config.Host, config.Port, config.Name))

	if er != nil {
		return nil, er
	}

	return &dashboardServiceImpl{
		connection: connection,
	}, nil
}

func (d dashboardServiceImpl) GetV1EarningsStats(ctx context.Context) ([]EarningStats, error) {
	stats := make([]EarningStats, 0, _statsNumberOfDays)

	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		return nil, errors.New("[dashboard service] user_id is not provided")
	}

	rows, er := d.connection.Query(ctx, "SELECT timestamp, amount, user_id, currency from public.earnings "+
		"WHERE user_id=$1", userID)
	if er != nil {
		return nil, er
	}

	for rows.Next() {
		var timestamp openapi_types.Date
		var amount float32

		if err := rows.Scan(&timestamp, &amount); err != nil {
			return nil, err
		}
		stats = append(stats, EarningStats{
			Date:     timestamp,
			Earnings: &amount,
		})
	}

	return stats, nil
}

func (d dashboardServiceImpl) GetV1Earnings(ctx context.Context) (float32, error) {
	var amount struct {
		value    float32
		currency string
	}

	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		return 0, errors.New("[dashboard service] user_id is not provided")
	}

	er := d.connection.QueryRow(ctx, "SELECT SUM(amount), currency from public.earnings WHERE user_id=$1 "+
		"group by user_id, currency", userID).Scan(&amount.value, &amount.currency)
	if er != nil {
		return 0, er
	}

	if amount.currency != _allowedCurrency {
		return 0, errors.New("[dashboard service] forbidden to use foreign currency, only one is allowed - " + _allowedCurrency)
	}

	return amount.value, nil
}
