package handler_test

import (
	"bootcamp/handler"
	"bootcamp/mock"
	"bootcamp/model"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetWalletDatas(t *testing.T) {

	t.Run("Should return all data by json format ", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockService := mock.NewMockIWalletService(mockCtrl)
		mockService.EXPECT().Wallets().Return(model.WalletsResponse{}, nil).Times(1)

		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		res := httptest.NewRecorder()

		walletHandler := handler.NewWalletHandler(mockService)
		walletHandler.GetWalletDatas(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		expectedResponse := model.WalletsResponse{}
		err := json.Unmarshal(res.Body.Bytes(), &expectedResponse)
		assert.Nil(t, err)
	})
	t.Run("should return server internal error when wallet service failed", func(t *testing.T) {
		mockService := mock.NewMockIWalletService(gomock.NewController(t))

		mockService.
			EXPECT().
			Wallets().
			Return(model.WalletsResponse{}, errors.New("error occured")). //with error
			Times(1)

		walletHandler := handler.NewWalletHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		res := httptest.NewRecorder()
		walletHandler.GetWalletDatas(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	})
}

func TestGetBalanceByUsername(t *testing.T) {

	t.Run("Should return correct balance data ", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockService := mock.NewMockIWalletService(mockCtrl)
		mockService.EXPECT().BalanceOfUser("testuser").Return(model.BalanceResponse{Balance: 1}, nil).Times(1)

		req := httptest.NewRequest(http.MethodGet, "/testuser", http.NoBody)
		res := httptest.NewRecorder()

		walletHandler := handler.NewWalletHandler(mockService)
		walletHandler.GetBalanceByUsername(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		expectedResponse := model.BalanceResponse{}
		err := json.Unmarshal(res.Body.Bytes(), &expectedResponse)

		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, model.BalanceResponse{Balance: 1})
	})
	t.Run("should return server internal error when wallet service failed", func(t *testing.T) {
		mockService := mock.NewMockIWalletService(gomock.NewController(t))
		mockService.
			EXPECT().BalanceOfUser("testuser").
			Return(model.BalanceResponse{}, errors.New("error occured")). //with error
			Times(1)

		walletHandler := handler.NewWalletHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/testuser", http.NoBody)
		res := httptest.NewRecorder()
		walletHandler.GetBalanceByUsername(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)

	})
}
func TestPutBalanceByUsername(t *testing.T) {

	t.Run("Should put correct data ", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockService := mock.NewMockIWalletService(mockCtrl)
		mockService.EXPECT().PutBalanceOfUser("testuser").Return(model.WalletResponse{Username: "testuser"}, nil).Times(1)

		req := httptest.NewRequest(http.MethodPut, "/testuser", http.NoBody)
		res := httptest.NewRecorder()

		walletHandler := handler.NewWalletHandler(mockService)
		walletHandler.PutBalanceByUsername(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		expectedResponse := model.WalletResponse{}
		err := json.Unmarshal(res.Body.Bytes(), &expectedResponse)

		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, model.WalletResponse{Username: "testuser", Balance: 0})
	})
	t.Run("should return server internal error when wallet service failed", func(t *testing.T) {
		mockService := mock.NewMockIWalletService(gomock.NewController(t))
		mockService.
			EXPECT().PutBalanceOfUser("testuser").
			Return(model.WalletResponse{}, errors.New("error occured")). //with error
			Times(1)

		walletHandler := handler.NewWalletHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/testuser", http.NoBody)
		res := httptest.NewRecorder()
		walletHandler.PutBalanceByUsername(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)

	})
}

func TestPostBalanceByUsername(t *testing.T) {

	t.Run("Should post correct data ", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockService := mock.NewMockIWalletService(mockCtrl)
		mockService.EXPECT().PostBalanceByUser("testuser", 1).Return(model.WalletResponse{Username: "testuser", Balance: 1}, nil).Times(1)
		jsonData := []byte(`{"balance":1}`)
		req := httptest.NewRequest(http.MethodPost, "/testuser", bytes.NewReader(jsonData))
		res := httptest.NewRecorder()

		walletHandler := handler.NewWalletHandler(mockService)
		walletHandler.PostBalanceByUsername(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		expectedResponse := model.WalletResponse{}
		err := json.Unmarshal(res.Body.Bytes(), &expectedResponse)

		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, model.WalletResponse{Username: "testuser", Balance: 1})
	})
	t.Run("Should post uncorrect data ", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockService := mock.NewMockIWalletService(mockCtrl)
		mockService.EXPECT().PostBalanceByUser("testuser", -101).Return(model.WalletResponse{}, nil).Times(1)
		jsonData := []byte(`{"balance":-101}`)
		req := httptest.NewRequest(http.MethodPost, "/testuser", bytes.NewReader(jsonData))
		res := httptest.NewRecorder()

		walletHandler := handler.NewWalletHandler(mockService)
		walletHandler.PostBalanceByUsername(res, req)

		assert.Equal(t, res.Result().StatusCode, http.StatusOK)
		expectedResponse := model.WalletResponse{}
		err := json.Unmarshal(res.Body.Bytes(), &expectedResponse)

		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, model.WalletResponse{})
	})
}
