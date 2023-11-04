-- name: CreateRenter :one
INSERT INTO renter (
    full_name
  ) VALUES (
      $1
) RETURNING *;


-- name: GetRenter :one
SELECT * FROM renter
WHERE id = $1 LIMIT 1;

-- name: GetRenters :many
SELECT * FROM renter
ORDER BY full_name
LIMIT $1
OFFSET $2;

/*DO NOT ALLOW CHANGE RENTER NAME*/

-- name: DeleteRenter :one
DELETE FROM renter
WHERE id = $1
RETURNING *;
