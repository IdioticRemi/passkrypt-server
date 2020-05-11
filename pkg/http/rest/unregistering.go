package rest

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/kizuru/passkrypt-server/pkg/unregistering"
	"net/http"
)

func unregisterAccount(s unregistering.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		decoder := json.NewDecoder(r.Body)

		var newUser unregistering.User
		err := decoder.Decode(&newUser)
		if err != nil {
			http.Error(w, "Invalid user object.", http.StatusBadRequest)
			return
		}

		err = s.UnregisterUser(newUser)
		if err == unregistering.ErrUserNotFound {
			http.Error(w, "The user does not exist.", http.StatusNotFound)
			return
		}

		_ = json.NewEncoder(w).Encode("Unregistered user.")
	}
}
