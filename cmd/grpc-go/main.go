package main

import (
	"log"
	"os"
	"strings"

	"github.com/Mussabaheen/GRPC_GO/config"
)

func main() {
	env := ""
	if envVar := os.Getenv("ENV"); envVar != "" {
		env = strings.ToLower(envVar)
	}
	if err := config.Load(env, "./config"); err != nil {
		log.Fatal("Could not load config:", err)
	}
}
