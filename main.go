package main

import (
	"fmt"
	"log"

	"github.com/abuelhassan/go-config-initialization/config"
	"github.com/abuelhassan/go-config-initialization/initializer"
	"github.com/abuelhassan/go-config-initialization/publisher"
)

func main() {
	cfg := config.Config{
		Port: 8080,
		Publisher: config.Publisher{
			Name: "Publisher",
		},
	}
	initializePackages(cfg)
	_ = publisher.New().Publish()
}

func initializePackages(cfg config.Config) {
	err := initializer.New().InitializePackages(cfg)
	if err != nil {
		panic(fmt.Errorf("config: %w", err))
		return
	}
	log.Println("All good!")
}
