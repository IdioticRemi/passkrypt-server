package rest

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/kizuru/passkrypt-server/pkg/registering"
	"net/http"
)

func registerUser(s registering.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		decoder := json.NewDecoder(r.Body)

		var newUser registering.User
		err := decoder.Decode(&newUser)
		if err != nil {
			http.Error(w, "Invalid user object.", http.StatusBadRequest)
			return
		}

		err = s.RegisterUser(newUser)
		if err == registering.ErrUserDuplicate {
			http.Error(w, "The user already exists.", http.StatusConflict)
			return
		}

		_ = json.NewEncoder(w).Encode("New user registered.")
	}
}
