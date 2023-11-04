-- name: CreateRentalAgreement :one
INSERT INTO rental_agreement (
  renter_id,
  rental_id,
  start_date,
  end_date,
  price
) VALUES (
$1, $2, $3, $4, $5
) RETURNING *;

-- name: GetRentalAgreement :one
SELECT * FROM rental_agreement 
WHERE id = $1 LIMIT 1;

-- name: GetRentalAgreementsByRenter :many
SELECT * FROM rental_agreement 
WHERE renter_id = $1 
ORDER BY rental_id
LIMIT $2
OFFSET $3;

-- name: ExtendRentalAgreement :one 
UPDATE rental_agreement
SET end_date = $1
WHERE id = $2
RETURNING *;

-- name: DeleteRentalAgreement :one
DELETE FROM rental_agreement
WHERE id = $1
RETURNING *;

