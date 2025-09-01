package main

import (
	"fmt"

	"github.com/dahakeadi15/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	registeredCommands map[string]func(s *state, cmd command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.registeredCommands[cmd.Name]
	if !exists {
		return fmt.Errorf("command '%s' does not exist", cmd.Name)
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
