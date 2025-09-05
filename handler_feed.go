package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dahakeadi15/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Arguments[0]
	url := cmd.Arguments[1]

	// add feed to db
	newFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    currUser.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(newFeed)
	fmt.Println("==========================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:        %s\n", feed.ID)
	fmt.Printf(" * Created:   %s\n", feed.CreatedAt)
	fmt.Printf(" * Updated:   %s\n", feed.UpdatedAt)
	fmt.Printf(" * Name:      %s\n", feed.Name)
	fmt.Printf(" * URL:       %s\n", feed.Url)
	fmt.Printf(" * UserId:    %s\n", feed.UserID)
}
