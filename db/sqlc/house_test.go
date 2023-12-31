package db

import (
	"context"
	"crypto/rand"
	"math/big"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TTNguyenDev/rental_track/util"
)

func TestCreateHouse(t *testing.T) {
	arg := GetRandomCreateHouseParams()
	house := CreateHouseForTest(arg)
	require.NotEmpty(t, house)

	require.Equal(t, arg.Name, house.Name)
	require.Equal(t, arg.Address, house.Address)
	require.Equal(t, arg.Kind, house.Kind)

	require.NotZero(t, house.ID)
	require.NotZero(t, house.CreatedAt)
}

func TestDeleteHouse(t *testing.T) {
	arg := GetRandomCreateHouseParams()
	createdHouse := CreateHouseForTest(arg)

	house, err := testQueries.DeleteHouse(context.Background(), createdHouse.ID)
	require.NoError(t, err)
	require.NotEmpty(t, house)

	require.Equal(t, createdHouse.ID, house.ID)
}

func TestGetHouse(t *testing.T) {
	arg := GetRandomCreateHouseParams()
	createdHouse := CreateHouseForTest(arg)

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

	arg1 := GetRandomCreateHouseParams()
	arg2 := GetRandomCreateHouseParams()
	createdHouses := []House{CreateHouseForTest(arg1), CreateHouseForTest(arg2)}

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
	_arg := GetRandomCreateHouseParams()
	createdHouse := CreateHouseForTest(_arg)

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
func CreateHouseForTest(arg CreateHouseParams) House {
	house, err := testQueries.CreateHouse(context.Background(), arg)
	if err != nil {
		panic(err)
	}
	return house
}

func GetRandomCreateHouseParams() CreateHouseParams {
	arg := CreateHouseParams{
		Name:    RandomName(),
		Address: RandomAddress(),
		Kind:    HousekindRooms,
	}
	return arg
}

func RandomName() string {
	return util.RandomString(6)
}

func RandomAddress() string {
	return util.RandomString(100)
}

func RandomPrice() string {
	return strconv.FormatInt(int64(util.RandomInt(0, 1000)), 10)
}

func RandomHouseKind() Housekind {
	kinds := []Housekind{HousekindHouse, HousekindRooms}
	n := int64(len(kinds))
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(n))
	if err != nil {
		panic(err)
	}
	return kinds[randomNumber.Int64()]
}

func RandomRentalStatus() Rentalstatus {
	statuses := []Rentalstatus{RentalstatusRented, RentalstatusEmpty}
	n := int64(len(statuses))

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(n))
	if err != nil {
		panic(err)
	}
	return statuses[randomNumber.Int64()]
}
