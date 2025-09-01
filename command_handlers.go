package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("username is required")
	}

	name := cmd.arguments[0]
	err := s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("user '%s' as been set\n", name)

	return nil
}
