-- name: FindAllCars :many
SELECT * FROM cars;

-- name: InsertCar :one
INSERT INTO cars (brand, model, year, price) VALUES (?, ?, ?, ?) RETURNING *;
