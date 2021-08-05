package inmemory

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mniak/Alkanoid/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionRepo_SaveWithoutIdAndLoad(t *testing.T) {
	var tra1 domain.Transaction
	gofakeit.Struct(&tra1)
	tra1.ID = 0

	sut := NewTransactionRepository()

	id, err := sut.Save(tra1)
	require.NoError(t, err)
	assert.Equal(t, 1, id)
	tra1.ID = id

	tra2, err := sut.Load(id)
	require.NoError(t, err)
	assert.Equal(t, id, tra2.ID)

	assert.Equal(t, tra1, tra2)
}

func TestTransactionRepo_LoadInvalid_ShouldReturnSpecificError(t *testing.T) {
	id := int(gofakeit.Int32())
	sut := NewTransactionRepository()

	_, err := sut.Load(id)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrNotFound, err)
}

func TestTransactionRepo_SaveBigId_ShouldUpdateInternalCounter(t *testing.T) {
	var tra500 domain.Transaction
	gofakeit.Struct(&tra500)
	tra500.ID = 500

	var traX domain.Transaction
	gofakeit.Struct(&traX)
	traX.ID = 0

	sut := NewTransactionRepository()

	id, err := sut.Save(tra500)
	require.NoError(t, err)
	assert.Equal(t, 500, id)

	id, err = sut.Save(traX)
	require.NoError(t, err)
	assert.Equal(t, 501, id)
}
