package main

import (
	"fmt"
	"menagment-app-2/src/config"
)

func main() {

	cfg, err := config.LoadConfigSingleton()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}
	fmt.Printf("Starting server on port %s\n", cfg.Port)
}
