package db 

import (
	"context"
)

//Querys to the database interface
type Querier interface {
	CreateCar(ctx context.Context, arg CreateCarParams) (Car, error)
	DeleteCar(ctx context.Context, id int32) error 
	GetCar(ctx context.Context, id int32) (Car, error) 
	GetCars(ctx context.Context) ([]Car, error)
	UpdateCar(ctx context.Context, arg UpdateCarParams) (Car, error)
}

var _Querier = (*Querier)(nil)