package repository

import (
	"bootcamp/model"
	"errors"
)

var W = model.WalletsResponse{
	{
		Username: "test1",
		Balance:  10,
	},
}

type IWalletRepository interface {
	GetAllWallets() (model.WalletsResponse, error)
	GetBalanceOfUser(username string) (model.BalanceResponse, error)
	PostWallet(username string, balance int) (model.WalletResponse, error)
	PutWallet(username string) (model.WalletResponse, error)
}

type WalletRepository struct {
	initialValue, minimumValue int
}

func (*WalletRepository) GetAllWallets() (model.WalletsResponse, error) {

	return W, nil
}

func (*WalletRepository) GetBalanceOfUser(username string) (model.BalanceResponse, error) {

	userBalance := model.BalanceResponse{}
Loop:
	for _, v := range W {
		if v.Username == username {
			userBalance.Balance = v.Balance
			break Loop
		}
	}

	return userBalance, nil
}

func (w *WalletRepository) PutWallet(username string) (model.WalletResponse, error) {

	newUser := model.WalletResponse{Username: username, Balance: w.initialValue}
	W = append(W, newUser)
	return newUser, nil
}
func (w *WalletRepository) PostWallet(username string, balance int) (model.WalletResponse, error) {

	newBalance := model.BalanceResponse{}
	findedIndex := 0
Loop:
	for i, v := range W {
		if v.Username == username {
			findedIndex = i
			newBalance.Balance = v.Balance + balance
			break Loop
		}
	}
	if newBalance.Balance < w.minimumValue {
		return model.WalletResponse{}, errors.New("Insufficient balance !")
	}
	W[findedIndex].Balance += balance
	return W[findedIndex], nil

}

func NewWalletRepository(initialValue, minimumValue int) IWalletRepository {
	return &WalletRepository{
		initialValue, minimumValue,
	}
}
