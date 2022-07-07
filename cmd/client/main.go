package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Mussabaheen/GRPC_GO/pkg/config"
	"github.com/Mussabaheen/GRPC_GO/pkg/greetings"
	"github.com/gorilla/mux"
)

func main() {
	env := ""
	if envVar := os.Getenv("ENV"); envVar != "" {
		env = strings.ToLower(envVar)
	}
	if err := config.Load(env, "../GRPC_GO/config"); err != nil {
		log.Fatal("Could not load config:", err)
	}
	r := mux.NewRouter()
	greetingsController := greetings.NewController(greetings.NewService())
	greetingsController.RegisterRoutes(r)
	port := config.Global.Port
	fmt.Println("Starting server :", port)
	err := http.ListenAndServe("localhost:3028", r)
	if err != nil {
		fmt.Println(err)
	}
}
