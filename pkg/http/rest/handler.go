package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Handler() http.Handler {
	router := httprouter.New()

	//router.GET("/credentials", getCredentials())

	//router.POST("/credentials", addCredentials())

	return router
}