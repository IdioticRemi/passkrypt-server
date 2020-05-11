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
		if err == listing.ErrAccountNotFound {
			http.Error(w, "The account does not exist.", http.StatusNotFound)
			return
		}

		_ = json.NewEncoder(w).Encode(account)
	}
}

func getUsers(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		list := s.GetUsers()

		_ = json.NewEncoder(w).Encode(list)
	}
}

func getUser(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		user, err := s.GetUser(p.ByName("id"))
		if err == listing.ErrUserNotFound {
			http.Error(w, "The user does not exist.", http.StatusNotFound)
			return
		}

		_ = json.NewEncoder(w).Encode(user)
	}
}
