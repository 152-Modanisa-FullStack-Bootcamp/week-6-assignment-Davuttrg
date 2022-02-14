package handler

import (
	"bootcamp/model"
	"bootcamp/service"
	"encoding/json"
	"net/http"
	"strings"
)

type IWalletHandler interface {
	GetWalletDatas(w http.ResponseWriter, r *http.Request)
	GetBalanceByUsername(w http.ResponseWriter, r *http.Request)
	PutBalanceByUsername(w http.ResponseWriter, r *http.Request)
	PostBalanceByUsername(w http.ResponseWriter, r *http.Request)
}

type WalletHandler struct {
	service service.IWalletService
}

func (ws *WalletHandler) GetWalletDatas(w http.ResponseWriter, r *http.Request) {
	response, err := ws.service.Wallets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	walletsData, _ := json.Marshal(response)
	//check response is json format
	w.Write(walletsData)
}
func (ws *WalletHandler) GetBalanceByUsername(w http.ResponseWriter, r *http.Request) {

	response, err := ws.service.BalanceOfUser(strings.Split(r.URL.Path, "/")[1])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	walletsData, _ := json.Marshal(response)
	w.Write(walletsData)
}
func (ws *WalletHandler) PutBalanceByUsername(w http.ResponseWriter, r *http.Request) {

	response, err := ws.service.PutBalanceOfUser(strings.Split(r.URL.Path, "/")[1])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	walletsData, _ := json.Marshal(response)
	w.Write(walletsData)
}

func (ws *WalletHandler) PostBalanceByUsername(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	reqBalance := &model.BalanceResponse{}
	err := decoder.Decode(reqBalance)
	if err != nil {
		panic(err)
	}

	response, err := ws.service.PostBalanceByUser(strings.Split(r.URL.Path, "/")[1], reqBalance.Balance)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

	walletsData, _ := json.Marshal(response)
	w.Write(walletsData)
}

func NewWalletHandler(service service.IWalletService) IWalletHandler {
	return &WalletHandler{
		service,
	}
}
