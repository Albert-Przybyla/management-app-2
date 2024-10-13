package main

import (
	"log"
	"menagment-app-2/src/api"
	"menagment-app-2/src/config"
)

func main() {

	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	a := api.New()
	a.Start()
}
