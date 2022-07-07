package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Mussabaheen/GRPC_GO/pkg/config"
	"github.com/Mussabaheen/GRPC_GO/pkg/greetings"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
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
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	greetingsService := greetings.NewService(*conn)
	greetingsController := greetings.NewController(greetingsService)
	greetingsController.RegisterRoutes(r)
	port := config.Global.Port
	fmt.Println("Starting server :", port)
	err = http.ListenAndServe("localhost:3028", r)
	if err != nil {
		log.Fatal(err)
	}

}
