package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateRentalAgreement(t *testing.T) RentalAgreement {
	rentalUnit := TestCreateRentalUnit(t)
	renter := TestCreateRenter(t)

	nextYear := time.Date(time.Now().Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
	arg := CreateRentalAgreementParams{
		RenterID:  renter.ID,
		RentalID:  rentalUnit.ID,
		StartDate: time.Now(),
		EndDate:   nextYear,
		Price:     RandomPrice(),
	}

	rentalAgreement, err := testQueries.CreateRentalAgreement(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rentalAgreement)

	require.Equal(t, arg.RenterID, rentalAgreement.RenterID)
	require.Equal(t, arg.RentalID, rentalAgreement.RentalID)
	require.Equal(t, arg.StartDate, rentalAgreement.StartDate)
	require.Equal(t, arg.StartDate, rentalAgreement.StartDate)
	require.Equal(t, arg.StartDate, rentalAgreement.StartDate)

	require.NotZero(t, rentalAgreement.ID)
	require.NotZero(t, rentalAgreement.CreatedAt)
	return rentalAgreement
}

func TestDeleteRentalAgreement(t *testing.T) {
	// createdHouse := TestCreateHouse(t)
	//
	// house, err := testQueries.DeleteHouse(context.Background(), createdHouse.ID)
	// require.NoError(t, err)
	// require.NotEmpty(t, house)
	//
	// require.Equal(t, createdHouse.ID, house.ID)
}

func TestGetRentalAgreement(t *testing.T) {
	// createdHouse := TestCreateHouse(t)
	//
	// house, err := testQueries.GetHouse(context.Background(), createdHouse.ID)
	// require.NoError(t, err)
	// require.NotEmpty(t, house)
	//
	// require.Equal(t, createdHouse.ID, house.ID)
	// require.Equal(t, createdHouse.Name, house.Name)
	// require.Equal(t, createdHouse.Address, house.Address)
	// require.Equal(t, createdHouse.Kind, house.Kind)
	// require.Equal(t, createdHouse.CreatedAt, house.CreatedAt)
}

func TestGetRentalAgreements(t *testing.T) {
	// // TODO: Delete all houses
	// createdHouses := []House{TestCreateHouse(t), TestCreateHouse(t)}
	//
	// arg := GetHousesParams{
	// 	Limit:  10,
	// 	Offset: 0,
	// }
	//
	// houses, err := testQueries.GetHouses(context.Background(), arg)
	// require.NoError(t, err)
	// require.NotEmpty(t, houses)
	//
	// require.Equal(t, len(houses), len(createdHouses))
	//
	// for i := 0; i < len(createdHouses); i++ {
	// 	require.Equal(t, houses[i], createdHouses[i])
	// }
}

func TestUpdateRentalAgreement(t *testing.T) {
	// createdHouse := TestCreateHouse(t)
	//
	// arg := UpdateHouseInfoParams{
	// 	Name:    RandomName(),
	// 	Address: RandomAddress(),
	// }
	//
	// house, err := testQueries.UpdateHouseInfo(context.Background(), arg)
	// require.NoError(t, err)
	// require.NotEmpty(t, house)
	//
	// require.Equal(t, house.ID, createdHouse.ID)
	// require.Equal(t, house.Name, arg.Name)
	// require.Equal(t, house.Address, arg.Address)
}
