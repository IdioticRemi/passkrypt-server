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
			http.Error(w, "Invalid account object.", http.StatusBadRequest)
			return
		}

		err = s.AddAccount(newAccount)
		if err == adding.ErrAccountDuplicate {
			http.Error(w, "The account already exists.", http.StatusConflict)
			return
		}

		_ = json.NewEncoder(w).Encode("New account added.")
	}
}
