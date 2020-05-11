package rest

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kizuru/passkrypt-server/pkg/adding"
	"github.com/kizuru/passkrypt-server/pkg/listing"
	"github.com/kizuru/passkrypt-server/pkg/logging"
	"github.com/kizuru/passkrypt-server/pkg/registering"
	"net/http"
)

func Handler(list listing.Service, add adding.Service, log logging.Service, reg registering.Service) http.Handler {
	router := httprouter.New()

	// Listing
	router.GET("/accounts", getAccounts(list))
	router.GET("/accounts/:id", getAccount(list))

	// Adding
	router.POST("/accounts", addAccount(add))

	// Auth/Logging

	// Auth/registering

	return router
}
