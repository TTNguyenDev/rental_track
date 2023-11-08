package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateRentalUnit(t *testing.T) RentalUnit {
	createdHouse := TestCreateHouse(t)
	arg := CreateRentalUnitParams{
		HouseID: createdHouse.ID,
		Price:   RandomPrice(),
		Status:  RandomRentalStatus(),
	}

	rentalUnit, err := testQueries.CreateRentalUnit(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rentalUnit)

	require.Equal(t, arg.HouseID, rentalUnit.HouseID)
	require.Equal(t, arg.Price, rentalUnit.Price)
	require.Equal(t, arg.Status, rentalUnit.Status)

	require.NotZero(t, rentalUnit.ID)
	require.NotZero(t, rentalUnit.UpdatedAt)
	return rentalUnit
}

func TestDeleteRentalUnit(t *testing.T) {
	createdRentalUnit := TestCreateRentalUnit(t)

	rentalUnit, err := testQueries.DeleteRentalUnit(context.Background(), createdRentalUnit.ID)
	require.NoError(t, err)
	require.NotEmpty(t, rentalUnit)

	require.Equal(t, createdRentalUnit.ID, rentalUnit.ID)
}

func TestGetRentalUnit(t *testing.T) {
	createdRentalUnit := TestCreateRentalUnit(t)

	rentalUnit, err := testQueries.GetRentalUnit(context.Background(), createdRentalUnit.ID)
	require.NoError(t, err)
	require.NotEmpty(t, rentalUnit)

	require.Equal(t, createdRentalUnit.ID, rentalUnit.ID)
	require.Equal(t, createdRentalUnit.HouseID, rentalUnit.HouseID)
	require.Equal(t, createdRentalUnit.Price, rentalUnit.Price)
	require.Equal(t, createdRentalUnit.Status, rentalUnit.Status)
	require.Equal(t, createdRentalUnit.UpdatedAt, rentalUnit.UpdatedAt)
}

func TestGetRentalUnits(t *testing.T) {
	// TODO: Delete all rentalUnits
	// createdRentalUnits := []RentalUnit{
	// 	TestCreateRentalUnit(t),
	// 	TestCreateRentalUnit(t),
	// 	TestCreateRentalUnit(t),
	// }

	arg := GetRentalUnitsByHouseParams{
		Limit:  10,
		Offset: 0,
	}

	rentalUnits, err := testQueries.GetRentalUnitsByHouse(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rentalUnits)

	//TODO CreateRentalUnitByHouseId
	// require.Equal(t, len(rentalUnits), len(createdRentalUnits))
	//
	// for i := 0; i < len(rentalUnits); i++ {
	// 	require.Equal(t, houses[i], createdHouses[i])
	// }
}

func TestUpdateRentalUnit(t *testing.T) {
	createdHouse := TestCreateHouse(t)

	arg := UpdateRentalUnitParams{
		ID:     createdHouse.ID,
		Price:  RandomPrice(),
		Status: RandomRentalStatus(),
	}

	rentalUnit, err := testQueries.UpdateRentalUnit(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rentalUnit)

	require.Equal(t, rentalUnit.HouseID, arg.ID)
	require.Equal(t, rentalUnit.Price, arg.Price)
	require.Equal(t, rentalUnit.Status, arg.Status)
}
