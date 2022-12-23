package main

import (
	"log"
	"net/http"

	"github.com/9996rojit/backend_go/pkg/router"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	router.HandleAuth(r)
	router.HandelPermissionRoute(r)
	log.Fatal(http.ListenAndServe("localhost:9910", r))
}