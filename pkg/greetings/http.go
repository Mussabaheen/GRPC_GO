package greetings

import (
	"net/http"

	httputil "github.com/Mussabaheen/GRPC_GO/pkg/httputil"

	"github.com/gorilla/mux"
)

type GretingsController struct {
	service Service
}

func NewController(service Service) *GretingsController {
	return &GretingsController{
		service: service,
	}
}

func (gc *GretingsController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/sayhello", gc.sayHelloEndpoint).Methods("GET")
}

func (gc *GretingsController) sayHelloEndpoint(rw http.ResponseWriter, r *http.Request) {
	resp, err := gc.service.SayHello()
	if err != nil {
		httputil.RespondWithError(rw, 400, err.Error())
	}
	httputil.RespondWithJSON(rw, 200, resp)
}
