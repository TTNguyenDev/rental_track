-- name: CreateHouse :one
INSERT INTO house (
  name,
  address,
  kind
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetHouse :one 
SELECT * FROM house 
WHERE id = $1 LIMIT 1;

-- name: GetHouses :many
SELECT * FROM house 
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: UpdateHouseInfo :one
UPDATE house 
SET name = $1, address = $2
WHERE id = $3 
RETURNING *;

/*
TODO: Delete Rental_unit, rental agreement before delete the house
 */
-- name: DeleteHouse :one
DELETE FROM house
WHERE id = $1
RETURNING *;
