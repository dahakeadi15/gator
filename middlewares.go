package main

import (
	"context"
	"fmt"

	"github.com/dahakeadi15/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't find current user: %w", err)
		}
		if err := handler(s, c, currUser); err != nil {
			return err
		}
		return nil
	}

}
