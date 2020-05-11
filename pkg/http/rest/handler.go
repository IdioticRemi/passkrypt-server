package rest

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kizuru/passkrypt-server/pkg/adding"
	"github.com/kizuru/passkrypt-server/pkg/listing"
	"github.com/kizuru/passkrypt-server/pkg/logging"
	"github.com/kizuru/passkrypt-server/pkg/registering"
	"github.com/kizuru/passkrypt-server/pkg/unregistering"
	"net/http"
)

func Handler(list listing.Service, add adding.Service, reg registering.Service, unreg unregistering.Service, log logging.Service) http.Handler {
	router := httprouter.New()

	// Listing
	router.GET("/accounts", getAccounts(list))
	router.GET("/accounts/:id", getAccount(list))

	router.GET("/users", getUsers(list))
	router.GET("/users/:id", getUser(list))

	// Adding
	router.POST("/accounts", addAccount(add))

	// Logging

	// Registering
	router.POST("/users", registerUser(reg))

	// Unregistering
	router.DELETE("/users", unregisterAccount(unreg))

	return router
}
