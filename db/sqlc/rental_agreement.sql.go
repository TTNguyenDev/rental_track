// Code generated by sqlc. DO NOT EDIT.
// source: rental_agreement.sql

package db

import (
	"context"
	"time"
)

const createRentalAgreement = `-- name: CreateRentalAgreement :one
INSERT INTO rental_agreement (
  renter_id,
  rental_id,
  start_date,
  end_date,
  price
) VALUES (
$1, $2, $3, $4, $5
) RETURNING id, renter_id, rental_id, start_date, end_date, price, created_at
`

type CreateRentalAgreementParams struct {
	RenterID  int32     `json:"renter_id"`
	RentalID  int32     `json:"rental_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Price     string    `json:"price"`
}

func (q *Queries) CreateRentalAgreement(ctx context.Context, arg CreateRentalAgreementParams) (RentalAgreement, error) {
	row := q.db.QueryRowContext(ctx, createRentalAgreement,
		arg.RenterID,
		arg.RentalID,
		arg.StartDate,
		arg.EndDate,
		arg.Price,
	)
	var i RentalAgreement
	err := row.Scan(
		&i.ID,
		&i.RenterID,
		&i.RentalID,
		&i.StartDate,
		&i.EndDate,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const deleteRentalAgreement = `-- name: DeleteRentalAgreement :one
DELETE FROM rental_agreement
WHERE id = $1
RETURNING id, renter_id, rental_id, start_date, end_date, price, created_at
`

func (q *Queries) DeleteRentalAgreement(ctx context.Context, id int32) (RentalAgreement, error) {
	row := q.db.QueryRowContext(ctx, deleteRentalAgreement, id)
	var i RentalAgreement
	err := row.Scan(
		&i.ID,
		&i.RenterID,
		&i.RentalID,
		&i.StartDate,
		&i.EndDate,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const extendRentalAgreement = `-- name: ExtendRentalAgreement :one
UPDATE rental_agreement
SET end_date = $1
WHERE id = $2
RETURNING id, renter_id, rental_id, start_date, end_date, price, created_at
`

type ExtendRentalAgreementParams struct {
	EndDate time.Time `json:"end_date"`
	ID      int32     `json:"id"`
}

func (q *Queries) ExtendRentalAgreement(ctx context.Context, arg ExtendRentalAgreementParams) (RentalAgreement, error) {
	row := q.db.QueryRowContext(ctx, extendRentalAgreement, arg.EndDate, arg.ID)
	var i RentalAgreement
	err := row.Scan(
		&i.ID,
		&i.RenterID,
		&i.RentalID,
		&i.StartDate,
		&i.EndDate,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const getRentalAgreement = `-- name: GetRentalAgreement :one
SELECT id, renter_id, rental_id, start_date, end_date, price, created_at FROM rental_agreement 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRentalAgreement(ctx context.Context, id int32) (RentalAgreement, error) {
	row := q.db.QueryRowContext(ctx, getRentalAgreement, id)
	var i RentalAgreement
	err := row.Scan(
		&i.ID,
		&i.RenterID,
		&i.RentalID,
		&i.StartDate,
		&i.EndDate,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const getRentalAgreementsByRenter = `-- name: GetRentalAgreementsByRenter :many
SELECT id, renter_id, rental_id, start_date, end_date, price, created_at FROM rental_agreement 
WHERE renter_id = $1 
ORDER BY rental_id
LIMIT $2
OFFSET $3
`

type GetRentalAgreementsByRenterParams struct {
	RenterID int32 `json:"renter_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) GetRentalAgreementsByRenter(ctx context.Context, arg GetRentalAgreementsByRenterParams) ([]RentalAgreement, error) {
	rows, err := q.db.QueryContext(ctx, getRentalAgreementsByRenter, arg.RenterID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RentalAgreement{}
	for rows.Next() {
		var i RentalAgreement
		if err := rows.Scan(
			&i.ID,
			&i.RenterID,
			&i.RentalID,
			&i.StartDate,
			&i.EndDate,
			&i.Price,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
