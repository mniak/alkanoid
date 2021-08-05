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

	requestObj := application.CreateAccountRequest{
		DocumentNumber: gofakeit.Numerify("##.###.###/####-##"),
	}
	responseObj := application.CreateAccountResponse{
		AccountID: gofakeit.Number(1, 100),
	}
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

	requestObj := application.GetAccountRequest{
		AccountID: gofakeit.Number(1, 100),
	}
	responseObj := application.GetAccountResponse{
		AccountID:      requestObj.AccountID,
		DocumentNumber: gofakeit.Numerify("##.###.###/####-##"),
	}
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
