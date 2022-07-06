package greetings

import (
	"net/http"

	"github.com/gorilla/mux"
)

type GretingsController struct {
	service Service
}

func NewController(service Service) *GretingsController {
	return &GretingsController{}
}

func (gc *GretingsController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/sayHello", gc.sayHelloEndpoint).Methods("GET")
}

func (gc *GretingsController) sayHelloEndpoint(rw http.ResponseWriter, r *http.Request) {
}
