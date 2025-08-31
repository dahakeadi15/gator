package main

import (
	"fmt"
	"log"

	"github.com/dahakeadi15/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("aditya")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
