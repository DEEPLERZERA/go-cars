package db

import (
	"context"
)

//Querys to the database
const createCar = `-- name: CreateCar :one
INSERT INTO cars (
	name,
	price,
	brand
) VALUES (
	$1, $2, $3
) RETURNING id, name, price, brand, created_at
`

type CreateCarParams struct {
	Name  string `json:"name"`
	Price int32  `json:"price"`
	Brand string `json:"brand"`
}

func (q *Queries) CreateCar(ctx context.Context, arg CreateCarParams) (Car, error) {
	row := q.db.QueryRowContext(ctx, createCar, arg.Name, arg.Price, arg.Brand)
	var car Car
	err := row.Scan(&car.ID, &car.Name, &car.Price, &car.Brand, &car.CreatedAt)
	return car, err
}

const deleteCar = `-- name: DeleteCar :exec
DELETE FROM cars
WHERE id = $1
`

func (q *Queries) DeleteCar(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCar, id)
	return err
}

const getCar = `-- name: GetCar :one
SELECT id, name, price, brand, created_at FROM cars
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCar(ctx context.Context, id int32) (Car, error) {
	row := q.db.QueryRowContext(ctx, getCar, id)
	var car Car
	err := row.Scan(&car.ID, &car.Name, &car.Price, &car.Brand, &car.CreatedAt)
	return car, err
}

const getCars = `-- name: GetCars :many
SELECT id, name, price, brand, created_at FROM cars
`

func (q *Queries) GetCars(ctx context.Context) ([]Car, error) {
	rows, err := q.db.QueryContext(ctx, getCars)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	carsList := []Car{}
	for rows.Next() {
		var car Car
		if err := rows.Scan(&car.ID, &car.Name, &car.Price, &car.Brand, &car.CreatedAt); err != nil {
			return nil, err
		}
		carsList = append(carsList, car)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return carsList, nil 
}

const updateCar = `-- name: UpdateCar :one
UPDATE cars 
SET name = $2, price = $3, brand = $4
WHERE id = $1 RETURNING id, name, price, brand, created_at
`

type UpdateCarParams struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
	Brand string `json:"brand"`
}

func (q *Queries) UpdateCar(ctx context.Context, arg UpdateCarParams) (Car, error) {
	row := q.db.QueryRowContext(ctx, updateCar, arg.ID, arg.Name, arg.Price, arg.Brand)
	var car Car
	err := row.Scan(&car.ID, &car.Name, &car.Price, &car.Brand, &car.CreatedAt)
	return car, err
}