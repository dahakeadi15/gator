package main

import (
	"fmt"

	"github.com/dahakeadi15/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
	}

	err = cfg.SetUser("aditya")
	if err != nil {
		fmt.Printf("error setting user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
	}
	fmt.Println("config:")
	fmt.Printf("  db_url: %s\n", cfg.DatabaseURL)
	fmt.Printf("  current_user_name: %s\n", cfg.CurrentUserName)
}
