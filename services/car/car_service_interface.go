package carservice

import (
	"context"

	"github.com/isaquecsilva/graphql/models"
)

type CarServiceInterface interface {
	CreateCar(ctx context.Context, request CreateCarRequest) (models.Car, error)
	GetAllCars(ctx context.Context) ([]models.Car, error)
}
