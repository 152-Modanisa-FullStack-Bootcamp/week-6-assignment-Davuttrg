package service

import (
	"bootcamp/model"
	"bootcamp/repository"
	"fmt"
)

type IWalletService interface {
	Wallets() (model.WalletsResponse, error)
	BalanceOfUser(username string) (model.BalanceResponse, error)
	PutBalanceOfUser(username string) (model.WalletResponse, error)
	PostBalanceByUser(username string, balance int) (model.WalletResponse, error)
}

type WalletService struct {
	repo repository.IWalletRepository
}

func (w *WalletService) Wallets() (model.WalletsResponse, error) {
	return w.repo.GetAllWallets()
}
func (w *WalletService) BalanceOfUser(username string) (model.BalanceResponse, error) {
	return w.repo.GetBalanceOfUser(username)
}
func (w *WalletService) PutBalanceOfUser(username string) (model.WalletResponse, error) {
	wallets, err := w.repo.GetAllWallets()
	if err != nil {
		return model.WalletResponse{}, err
	}
	existUser := model.WalletsResponse{}
Loop:
	for _, v := range wallets {
		if v.Username == username {
			existUser = append(existUser, v)
			break Loop
		}
	}

	if length := len(existUser); length > 0 {
		return existUser[0], nil
	} else {
		return w.repo.PutWallet(username)
	}

}
func (w *WalletService) PostBalanceByUser(username string, balance int) (model.WalletResponse, error) {
	fmt.Println("username :", username)
	fmt.Println("balance :", balance)
	return w.repo.PostWallet(username, balance)
}
func NewWalletService(repo repository.IWalletRepository) IWalletService {
	return &WalletService{
		repo: repo,
	}
}
