package service_test

import (
	"bootcamp/mock"
	"bootcamp/model"
	"bootcamp/service"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestWallets(t *testing.T) {
	//get mocked interface
	mockCtrl := gomock.NewController(t)
	mockObj := mock.NewMockIWalletRepository(mockCtrl)

	//writed expectation for 1 time
	mockObj.EXPECT().GetAllWallets().Return(model.WalletsResponse{}, nil).Times(1)

	//inject IWalletRepository dependency to WalletService
	walletService := service.NewWalletService(mockObj)

	wallets, err := walletService.Wallets()
	assert.Equal(t, model.WalletsResponse{}, wallets)
	assert.Nil(t, err)
}
func TestBalanceOfUser(t *testing.T) {
	//get mocked interface
	mockCtrl := gomock.NewController(t)
	mockObj := mock.NewMockIWalletRepository(mockCtrl)

	//writed expectation for 1 time
	mockObj.EXPECT().GetBalanceOfUser("testuser").Return(model.BalanceResponse{Balance: 1}, nil).Times(1)

	//inject IWalletRepository dependency to WalletService
	walletService := service.NewWalletService(mockObj)

	balance, err := walletService.BalanceOfUser("testuser")
	assert.Equal(t, 1, balance.Balance)
	assert.Nil(t, err)
}

func TestPutBalanceOfUser(t *testing.T) {
	t.Run("put operation for non exist user", func(t *testing.T) {
		//get mocked interface
		mockCtrl := gomock.NewController(t)
		mockObj := mock.NewMockIWalletRepository(mockCtrl)

		//writed expectation for 1 time

		mockObj.EXPECT().GetAllWallets().Return(model.WalletsResponse{{Username: "testuser", Balance: 0}}, nil).Times(1)
		mockObj.EXPECT().PutWallet("testuser1").Return(model.WalletResponse{Username: "testuser1", Balance: 0}, nil).Times(1)

		//inject IWalletRepository dependency to WalletService
		walletService := service.NewWalletService(mockObj)

		user, err := walletService.PutBalanceOfUser("testuser1")
		assert.Equal(t, model.WalletResponse{Username: "testuser1", Balance: 0}, user)
		assert.Nil(t, err)
	})
	t.Run("put operation for exist user", func(t *testing.T) {
		//get mocked interface
		mockCtrl := gomock.NewController(t)
		mockObj := mock.NewMockIWalletRepository(mockCtrl)

		//writed expectation for 1 time

		mockObj.EXPECT().GetAllWallets().Return(model.WalletsResponse{{Username: "testuser", Balance: 0}}, nil).Times(1)
		mockObj.EXPECT().PutWallet("testuser").Return(model.WalletResponse{Username: "testuser", Balance: 0}, nil).Times(0)

		//inject IWalletRepository dependency to WalletService
		walletService := service.NewWalletService(mockObj)

		user, err := walletService.PutBalanceOfUser("testuser")
		assert.Equal(t, model.WalletResponse{Username: "testuser", Balance: 0}, user)
		assert.Nil(t, err)
	})
}
func TestPostBalanceOfUser(t *testing.T) {
	t.Run("post operation for valid body", func(t *testing.T) {
		//get mocked interface
		mockCtrl := gomock.NewController(t)
		mockObj := mock.NewMockIWalletRepository(mockCtrl)

		//writed expectation for 1 time
		mockObj.EXPECT().PostWallet("testuser", 1).Return(model.WalletResponse{Username: "testuser", Balance: 1}, nil).Times(1)

		//inject IWalletRepository dependency to WalletService
		walletService := service.NewWalletService(mockObj)

		user, err := walletService.PostBalanceByUser("testuser", 1)
		assert.Equal(t, model.WalletResponse{Username: "testuser", Balance: 1}, user)
		assert.Nil(t, err)
	})
	t.Run("post operation for non valid body", func(t *testing.T) {
		//get mocked interface
		mockCtrl := gomock.NewController(t)
		mockObj := mock.NewMockIWalletRepository(mockCtrl)

		//writed expectation for 1 time
		mockObj.EXPECT().PostWallet("testuser", -101).Return(model.WalletResponse{}, errors.New("Insufficient balance !")).Times(1)

		//inject IWalletRepository dependency to WalletService
		walletService := service.NewWalletService(mockObj)

		user, err := walletService.PostBalanceByUser("testuser", -101)
		assert.Equal(t, model.WalletResponse{}, user)
		assert.NotNil(t, err)
	})
}
