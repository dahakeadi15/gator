package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dahakeadi15/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Arguments[0]

	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return fmt.Errorf("user does not exist")
		}
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Arguments[0]

	// create new user
	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return fmt.Errorf("user with that name already exists")
		}
		return fmt.Errorf("error creating user: %+v", err)
	}

	// set new user to current user
	err = s.cfg.SetUser(newUser.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	printUser(newUser)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list all users: %w", err)
	}

	for _, u := range users {
		name := u.Name
		if s.cfg.CurrentUserName == name {
			name += " (current)"
		}
		fmt.Printf(" * %s\n", name)
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:     %v\n", user.ID)
	fmt.Printf(" * Name:   %v\n", user.Name)
}
