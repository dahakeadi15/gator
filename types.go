package main

import "github.com/dahakeadi15/gator/internal/config"

type state struct {
	cfg *config.Config
}

type command struct {
	name      string
	arguments []string
}
