// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Housekind string

const (
	HousekindHouse Housekind = "House"
	HousekindRooms Housekind = "Rooms"
)

func (e *Housekind) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Housekind(s)
	case string:
		*e = Housekind(s)
	default:
		return fmt.Errorf("unsupported scan type for Housekind: %T", src)
	}
	return nil
}

type Rentalstatus string

const (
	RentalstatusRented Rentalstatus = "Rented"
	RentalstatusEmpty  Rentalstatus = "Empty"
)

func (e *Rentalstatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Rentalstatus(s)
	case string:
		*e = Rentalstatus(s)
	default:
		return fmt.Errorf("unsupported scan type for Rentalstatus: %T", src)
	}
	return nil
}

type House struct {
	ID        int32        `json:"id"`
	Name      string       `json:"name"`
	Address   string       `json:"address"`
	Kind      Housekind    `json:"kind"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type RentalAgreement struct {
	ID        int32          `json:"id"`
	RenterID  int32          `json:"renter_id"`
	RentalID  int32          `json:"rental_id"`
	StartDate time.Time      `json:"start_date"`
	EndDate   sql.NullTime   `json:"end_date"`
	Price     sql.NullString `json:"price"`
	CreatedAt time.Time      `json:"created_at"`
}

type RentalUnit struct {
	ID        int32        `json:"id"`
	HouseID   int32        `json:"house_id"`
	Price     string       `json:"price"`
	Status    Rentalstatus `json:"status"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Renter struct {
	ID       int32  `json:"id"`
	FullName string `json:"full_name"`
}
