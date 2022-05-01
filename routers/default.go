package routers

import (
	"log"
	"net/http"
	"os"

	"github.com/Gprisco/decanto-golang/handlers"
	"github.com/gorilla/mux"
)

func Config(router *mux.Router) {
	logger := log.New(os.Stdout, "router (default) ", log.LstdFlags)

	getRouter := router.Methods(http.MethodPost).Subrouter()

	helloHandler := handlers.NewHello(logger)

	getRouter.HandleFunc("/", helloHandler.HandleGet)
}
