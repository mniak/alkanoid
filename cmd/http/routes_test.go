package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/Alkanoid/application"
	"github.com/mniak/Alkanoid/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var requestObj application.CreateAccountRequest
	var responseObj application.CreateAccountResponse
	gofakeit.Struct(&requestObj)
	gofakeit.Struct(&responseObj)

	app := mocks.NewMockApplication(ctrl)
	app.EXPECT().
		CreateAccount(requestObj).
		Return(responseObj, nil)

	router := setupRouter(app)

	w := httptest.NewRecorder()

	reqBytes, err := json.Marshal(requestObj)
	require.NoError(t, err)

	req, _ := http.NewRequest("POST", "/accounts", bytes.NewReader(reqBytes))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	var parsedResponse application.CreateAccountResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &parsedResponse))

	require.Equal(t, responseObj, parsedResponse)
}

func TestGetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var requestObj application.GetAccountRequest
	var responseObj application.GetAccountResponse
	gofakeit.Struct(&requestObj)
	gofakeit.Struct(&responseObj)

	app := mocks.NewMockApplication(ctrl)
	app.EXPECT().
		GetAccount(requestObj).
		Return(responseObj, nil)

	router := setupRouter(app)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/accounts/%d", requestObj.AccountID), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var parsedResponse application.GetAccountResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &parsedResponse))

	require.Equal(t, responseObj, parsedResponse)
}

func TestCreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var requestObj application.CreateTransactionRequest
	var responseObj application.CreateTransactionResponse
	gofakeit.Struct(&requestObj)
	gofakeit.Struct(&responseObj)

	app := mocks.NewMockApplication(ctrl)
	app.EXPECT().
		CreateTransaction(requestObj).
		Return(responseObj, nil)

	router := setupRouter(app)

	w := httptest.NewRecorder()

	reqBytes, err := json.Marshal(requestObj)
	require.NoError(t, err)

	req, _ := http.NewRequest("POST", "/transactions", bytes.NewReader(reqBytes))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	var parsedResponse application.CreateTransactionResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &parsedResponse))

	require.Equal(t, responseObj, parsedResponse)
}
