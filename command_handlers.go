package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("expected 1 argument, none was given")
	}

	name := cmd.arguments[0]
	err := s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("user '%s' as been set", name)

	return nil
}
