package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateReneterForTest(arg string) Renter {
	renter, err := testQueries.CreateRenter(context.Background(), arg)
	if err != nil {
		panic(err)
	}
	return renter
}

func TestCreateRenter(t *testing.T) {
	arg := RandomName()
	renter := CreateReneterForTest(arg)
	require.NotEmpty(t, renter)

	require.Equal(t, arg, renter.FullName)

	require.NotZero(t, renter.ID)
}

func TestDeleteRenter(t *testing.T) {
	arg := RandomName()
	createdRenter := CreateReneterForTest(arg)

	renter, err := testQueries.DeleteRenter(context.Background(), createdRenter.ID)
	require.NoError(t, err)
	require.NotEmpty(t, renter)

	require.Equal(t, createdRenter.ID, renter.ID)
}

func TestGetRenter(t *testing.T) {
	arg := RandomName()
	createdRenter := CreateReneterForTest(arg)

	renter, err := testQueries.GetRenter(context.Background(), createdRenter.ID)
	require.NoError(t, err)
	require.NotEmpty(t, renter)

	require.Equal(t, createdRenter.ID, renter.ID)
	require.Equal(t, createdRenter.FullName, renter.FullName)
}

func TestGetRenters(t *testing.T) {
	// TODO: Delete all houses
	createdRenters := []Renter{
		CreateReneterForTest(RandomName()),
		CreateReneterForTest(RandomName()),
	}

	arg := GetRentersParams{
		Limit:  10,
		Offset: 0,
	}

	renters, err := testQueries.GetRenters(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, renters)

	require.Equal(t, len(renters), len(createdRenters))

	for i := 0; i < len(createdRenters); i++ {
		require.Equal(t, renters[i], createdRenters[i])
	}
}
