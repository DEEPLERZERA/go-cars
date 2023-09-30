-- name: CreateCar :one
INSERT INTO cars (
    name,
    price,
    brand
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetCar :one
SELECT * FROM cars
WHERE id = $1 LIMIT 1;

-- name: GetCars :many
SELECT * FROM cars;

-- name: UpdateCar :one
UPDATE cars SET
    name = $2,
    price = $3,
    brand = $4
WHERE id = $1 RETURNING *;

-- name: DeleteCar :exec
DELETE FROM cars
WHERE id = $1;



