package rest

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/kizuru/passkrypt-server/pkg/adding"
	"net/http"
)

func addAccount(s adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		decoder := json.NewDecoder(r.Body)

		var newAccount adding.Account
		err := decoder.Decode(&newAccount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode("Invalid account object.")
			return
		}

		err = s.AddAccount(newAccount)
		if err == adding.ErrDuplicate {
			w.WriteHeader(http.StatusConflict)
			_ = json.NewEncoder(w).Encode("The account already exists.")
			return
		}

		_ = json.NewEncoder(w).Encode("New Account added.")
	}
}
