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
	"github.com/mniak/Alkanoid/app"
	"github.com/mniak/Alkanoid/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var requestObj app.CreateAccountRequest
	var responseObj app.CreateAccountResponse
	gofakeit.Struct(&requestObj)
	gofakeit.Struct(&responseObj)

	a := mocks.NewMockApplication(ctrl)
	a.EXPECT().
		CreateAccount(requestObj).
		Return(responseObj, nil)

	router := setupRouter(a)

	w := httptest.NewRecorder()

	reqBytes, err := json.Marshal(requestObj)
	require.NoError(t, err)

	req, _ := http.NewRequest("POST", "/accounts", bytes.NewReader(reqBytes))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	var parsedResponse app.CreateAccountResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &parsedResponse))

	require.Equal(t, responseObj, parsedResponse)
}

func TestGetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var requestObj app.GetAccountRequest
	var responseObj app.GetAccountResponse
	gofakeit.Struct(&requestObj)
	gofakeit.Struct(&responseObj)

	a := mocks.NewMockApplication(ctrl)
	a.EXPECT().
		GetAccount(requestObj).
		Return(responseObj, nil)

	router := setupRouter(a)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/accounts/%d", requestObj.AccountID), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var parsedResponse app.GetAccountResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &parsedResponse))

	require.Equal(t, responseObj, parsedResponse)
}

func TestCreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var requestObj app.CreateTransactionRequest
	var responseObj app.CreateTransactionResponse
	gofakeit.Struct(&requestObj)
	gofakeit.Struct(&responseObj)

	a := mocks.NewMockApplication(ctrl)
	a.EXPECT().
		CreateTransaction(requestObj).
		Return(responseObj, nil)

	router := setupRouter(a)

	w := httptest.NewRecorder()

	reqBytes, err := json.Marshal(requestObj)
	require.NoError(t, err)

	req, _ := http.NewRequest("POST", "/transactions", bytes.NewReader(reqBytes))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	var parsedResponse app.CreateTransactionResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &parsedResponse))

	require.Equal(t, responseObj, parsedResponse)
}
