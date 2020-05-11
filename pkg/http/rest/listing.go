package rest

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/kizuru/passkrypt-server/pkg/listing"
	"net/http"
)

func getAccounts(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		list := s.GetAccounts()

		_ = json.NewEncoder(w).Encode(list)
	}
}

func getAccount(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		account, err := s.GetAccount(p.ByName("id"))
		if err == listing.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode("The account does not exist.")
			return
		}

		_ = json.NewEncoder(w).Encode(account)
	}
}
