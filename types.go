package main

import (
	"fmt"

	"github.com/dahakeadi15/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	functions map[string]func(s *state, cmd command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.functions[cmd.name]
	if !exists {
		return fmt.Errorf("command '%s' does not exist", cmd.name)
	}

	err := handler(s, cmd)
	if err != nil {
		return fmt.Errorf("error while running command '%s': %w", cmd.name, err)
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.functions[name] = f
}
