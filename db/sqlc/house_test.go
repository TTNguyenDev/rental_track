package db

import (
	"context"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TTNguyenDev/rental_track/util"
)

func TestCreateHouse(t *testing.T) House {
	arg := CreateHouseParams{
		Name:    RandomName(),
		Address: RandomAddress(),
		Kind:    HousekindRooms,
	}

	house, err := testQueries.CreateHouse(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, house)

	require.Equal(t, arg.Name, house.Name)
	require.Equal(t, arg.Address, house.Address)
	require.Equal(t, arg.Kind, house.Kind)

	require.NotZero(t, house.ID)
	require.NotZero(t, house.CreatedAt)
	return house
}

func TestDeleteHouse(t *testing.T) {
	createdHouse := TestCreateHouse(t)

	house, err := testQueries.DeleteHouse(context.Background(), createdHouse.ID)
	require.NoError(t, err)
	require.NotEmpty(t, house)

	require.Equal(t, createdHouse.ID, house.ID)
}

func TestGetHouse(t *testing.T) {
	createdHouse := TestCreateHouse(t)

	house, err := testQueries.GetHouse(context.Background(), createdHouse.ID)
	require.NoError(t, err)
	require.NotEmpty(t, house)

	require.Equal(t, createdHouse.ID, house.ID)
	require.Equal(t, createdHouse.Name, house.Name)
	require.Equal(t, createdHouse.Address, house.Address)
	require.Equal(t, createdHouse.Kind, house.Kind)
	require.Equal(t, createdHouse.CreatedAt, house.CreatedAt)
}

func TestGetHouses(t *testing.T) {
	// TODO: Delete all houses
	createdHouses := []House{TestCreateHouse(t), TestCreateHouse(t)}

	arg := GetHousesParams{
		Limit:  10,
		Offset: 0,
	}

	houses, err := testQueries.GetHouses(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, houses)

	require.Equal(t, len(houses), len(createdHouses))

	for i := 0; i < len(createdHouses); i++ {
		require.Equal(t, houses[i], createdHouses[i])
	}
}

func TestUpdateHouse(t *testing.T) {
	createdHouse := TestCreateHouse(t)

	arg := UpdateHouseInfoParams{
		Name:    RandomName(),
		Address: RandomAddress(),
	}

	house, err := testQueries.UpdateHouseInfo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, house)

	require.Equal(t, house.ID, createdHouse.ID)
	require.Equal(t, house.Name, arg.Name)
	require.Equal(t, house.Address, arg.Address)
}

// Util
func RandomName() string {
	return util.RandomString(6)
}

func RandomAddress() string {
	return util.RandomString(100)
}

func RandomPrice() int64 {
	return util.RandomInt(0, 1000)
}

func RandomHouseKind() Housekind {
	kinds := []Housekind{HousekindHouse, HousekindRooms}
	n := len(kinds)
	return kinds[rand.Intn(n)]
}
