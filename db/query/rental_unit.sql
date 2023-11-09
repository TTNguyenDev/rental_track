-- name: CreateRentalUnit :one
INSERT INTO rental_unit (
  house_id,
  price,
  status
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetRentalUnit :one 
SELECT * FROM rental_unit 
WHERE id = $1 LIMIT 1;

-- name: GetRentalUnitsByHouse :many
SELECT * FROM rental_unit 
WHERE house_id = $1
ORDER BY price
LIMIT $2
OFFSET $3;

-- name: UpdateRentalUnit :one
UPDATE rental_unit
SET price = $1, status = $2
WHERE id = $3
RETURNING *;

-- name: DeleteRentalUnit :one 
DELETE FROM rental_unit 
WHERE id = $1
RETURNING *;

-- name: DeleteRentalUnitsByHouse :many
DELETE FROM rental_unit
WHERE house_id = $1
RETURNING id;
