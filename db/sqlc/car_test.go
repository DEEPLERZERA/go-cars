package db 

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

//Test the creation of a car
func createRandomCar(t *testing.T) Car {
	arg := CreateCarParams{
		Name:  "BMW X5",
		Price:  60000,
		Brand:  "BMW",
	}

	car, err := testQueries.CreateCar(context.Background(), arg)


	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, arg.Name, car.Name)
	require.Equal(t, arg.Price, car.Price)
	require.Equal(t, arg.Brand, car.Brand)

	require.NotEmpty(t, car.ID)
	require.NotEmpty(t, car.CreatedAt)

	return car
}

func TestCreateCar(t *testing.T) {
	createRandomCar(t)
}

//Test the search of a car
func TestGetCar(t *testing.T) {
	carRandomCreated := createRandomCar(t)
	carFinded, err := testQueries.GetCar(context.Background(), carRandomCreated.ID)

	require.NoError(t, err)
	require.NotEmpty(t, carRandomCreated)
	require.NotEmpty(t, carFinded)

	require.Equal(t, carRandomCreated.ID, carFinded.ID)
	require.Equal(t, carRandomCreated.Name, carFinded.Name)
	require.Equal(t, carRandomCreated.Price, carFinded.Price)
	require.Equal(t, carRandomCreated.Brand, carFinded.Brand)
	require.Equal(t, carRandomCreated.CreatedAt, carFinded.CreatedAt)
}

//Test the deletion of a car
func TestDeleteCar(t *testing.T) {
	carRandomCreated := createRandomCar(t)
	err := testQueries.DeleteCar(context.Background(), carRandomCreated.ID)

	require.NoError(t, err)
}

//Test the update of a car
func TestUpdateCar(t *testing.T) {
	carRandomCreated := createRandomCar(t)

	arg := UpdateCarParams{
		ID: carRandomCreated.ID,
		Name: "BMW X6",
		Price:  70000,
		Brand:  "BMW",
	}

	carUpdated, err := testQueries.UpdateCar(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, carUpdated)

	require.Equal(t, carRandomCreated.ID, carUpdated.ID)
	require.Equal(t, arg.Name, carUpdated.Name)
	require.Equal(t, arg.Price, carUpdated.Price)
	require.Equal(t, arg.Brand, carUpdated.Brand)
	require.Equal(t, carRandomCreated.CreatedAt, carUpdated.CreatedAt)
}

//Test the search of all cars
func TestGetCars(t *testing.T) {
	carRandomCreated := createRandomCar(t)
	carsFinded, err := testQueries.GetCars(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, carRandomCreated)
	require.NotEmpty(t, carsFinded)

	for _, car := range carsFinded {
		require.NotEmpty(t, car.ID)
		require.NotEmpty(t, car.Name)
		require.NotEmpty(t, car.Price)
		require.NotEmpty(t, car.Brand)
		require.NotEmpty(t, car.CreatedAt)
	}
}