package carservice

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/isaquecsilva/graphql/models"
)

type CarServiceImpl struct {
	querier models.Querier
}

func (c *CarServiceImpl) CreateCar(ctx context.Context, request CreateCarRequest) (models.Car, error) {
	car, err := c.querier.InsertCar(ctx, models.InsertCarParams{
		Brand: request.Brand,
		Model: request.Model,
		Year:  request.Year,
		Price: request.Price,
	})

	if err != nil {
		slog.Error("creating new car", slog.Any("error", err), slog.Any("request", request))
		return models.Car{}, fmt.Errorf("error creating new car")
	}

	return car, nil
}

func (c *CarServiceImpl) GetAllCars(ctx context.Context) ([]models.Car, error) {
	cars, err := c.querier.FindAllCars(ctx)
	if err != nil {
		slog.Error("retrieving all cars", slog.Any("error", err))
		return nil, fmt.Errorf("error retrieving cars")
	}

	return cars, nil
}

func NewCarServiceImpl(q models.Querier) CarServiceInterface {
	return &CarServiceImpl{
		querier: q,
	}
}
