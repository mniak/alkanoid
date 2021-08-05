package inmemory

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mniak/Alkanoid/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountRepo_SaveWithoutIdAndLoad(t *testing.T) {
	var acc1 domain.Account
	gofakeit.Struct(&acc1)
	acc1.ID = 0

	sut := NewAccountRepository()

	id, err := sut.Save(acc1)
	require.NoError(t, err)
	assert.Equal(t, 1, id)
	acc1.ID = id

	acc2, err := sut.Load(id)
	require.NoError(t, err)
	assert.Equal(t, id, acc2.ID)

	assert.Equal(t, acc1, acc2)
}

func TestAccountRepo_LoadInvalid_ShouldReturnSpecificError(t *testing.T) {
	id := int(gofakeit.Int32())
	sut := NewAccountRepository()

	_, err := sut.Load(id)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrAccountNotFound, err)
}

func TestAccountRepo_SaveBigId_ShouldUpdateInternalCounter(t *testing.T) {
	var acc500 domain.Account
	gofakeit.Struct(&acc500)
	acc500.ID = 500

	var accX domain.Account
	gofakeit.Struct(&accX)
	accX.ID = 0

	sut := NewAccountRepository()

	id, err := sut.Save(acc500)
	require.NoError(t, err)
	assert.Equal(t, 500, id)

	id, err = sut.Save(accX)
	require.NoError(t, err)
	assert.Equal(t, 501, id)
}
