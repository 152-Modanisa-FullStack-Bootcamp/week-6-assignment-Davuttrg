package main

import (
	"bootcamp/config"
	"bootcamp/handler"
	"bootcamp/repository"
	"bootcamp/service"
	"net/http"
	"regexp"
)

var (
	listWallets = regexp.MustCompile(`^\/[\/]*$`)
	withParam   = regexp.MustCompile(`^\/[a-zA-Z0-9]*$`)
)

type walletsHandler struct{}

func (h *walletsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	conf := *config.Get()
	walletRepository := repository.NewWalletRepository(conf.InitialBalanceAmount, conf.MinumumBalanceAmount)
	walletService := service.NewWalletService(walletRepository)
	walletHandler := handler.NewWalletHandler(walletService)
	switch {
	case r.Method == http.MethodGet && listWallets.MatchString(r.URL.Path):
		walletHandler.GetWalletDatas(w, r)
		return
	case r.Method == http.MethodGet && withParam.MatchString(r.URL.Path):
		walletHandler.GetBalanceByUsername(w, r)
		return
	case r.Method == http.MethodPut && withParam.MatchString(r.URL.Path):
		walletHandler.PutBalanceByUsername(w, r)
		return
	case r.Method == http.MethodPost && withParam.MatchString(r.URL.Path):
		walletHandler.PostBalanceByUsername(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &walletsHandler{})
	http.ListenAndServe(":8080", mux)
}
