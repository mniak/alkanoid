package main

import (
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
	err := json.Unmarshal(w.Body.Bytes(), &parsedResponse)
	require.NoError(t, err)

	require.Equal(t, responseObj, parsedResponse)
}
