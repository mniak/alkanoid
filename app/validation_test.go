package app_test

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/Alkanoid/app"
	"github.com/mniak/Alkanoid/domain"
	"github.com/mniak/Alkanoid/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionRepoWithValidation_Save_WhenIsValid_ShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var trans domain.Transaction
	gofakeit.Struct(&trans)
	expectedId := int(gofakeit.Int32())

	inner := mocks.NewMockTransactionRepository(ctrl)
	inner.EXPECT().Save(trans).Return(expectedId, nil)

	valsvc := mocks.NewMockTransactionValidationService(ctrl)
	valsvc.EXPECT().Validate(trans).Return(domain.ValidResult())

	sut := app.WrapTransactionRepoWithValidation(inner, valsvc)

	id, err := sut.Save(trans)
	require.NoError(t, err)
	assert.Equal(t, expectedId, id)
}

func TestTransactionRepoWithValidation_Save_WhenHasMessage_ShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var trans domain.Transaction
	gofakeit.Struct(&trans)

	msg := gofakeit.Sentence(5)
	inner := mocks.NewMockTransactionRepository(ctrl)
	valsvc := mocks.NewMockTransactionValidationService(ctrl)
	valsvc.EXPECT().Validate(trans).Return(domain.InvalidResult(msg))

	sut := app.WrapTransactionRepoWithValidation(inner, valsvc)

	_, err := sut.Save(trans)
	require.Error(t, err)
	assert.Contains(t, err.Error(), msg)
}

func TestTransactionRepoWithValidation_Save_WhenHasError_ShouldReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var trans domain.Transaction
	gofakeit.Struct(&trans)

	expectedError := errors.New(gofakeit.Sentence(5))

	inner := mocks.NewMockTransactionRepository(ctrl)
	valsvc := mocks.NewMockTransactionValidationService(ctrl)
	valsvc.EXPECT().Validate(trans).Return(domain.ValidationErrorResult(expectedError))

	sut := app.WrapTransactionRepoWithValidation(inner, valsvc)

	_, err := sut.Save(trans)
	require.Error(t, err)
	assert.Contains(t, err.Error(), expectedError.Error())
}
