package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kizuru/passkrypt-server/pkg/http/rest"
)

func main() {
	router := rest.Handler()

	fmt.Println("The passkrypt server is now running: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
