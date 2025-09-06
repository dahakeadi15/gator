package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dahakeadi15/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Arguments[0]
	url := cmd.Arguments[1]

	// add feed to db
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed, user)
	fmt.Println()

	// add feed to current user's following
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed followed successfully.")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("===========================")

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list all feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for i, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user by user.id: %w", err)
		}

		fmt.Printf("%d:\n", i+1)
		printFeed(feed, user)
	}
	fmt.Println("=================")

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf(" * ID:        %s\n", feed.ID)
	fmt.Printf(" * Created:   %s\n", feed.CreatedAt)
	fmt.Printf(" * Updated:   %s\n", feed.UpdatedAt)
	fmt.Printf(" * Name:      %s\n", feed.Name)
	fmt.Printf(" * URL:       %s\n", feed.Url)
	fmt.Printf(" * User:      %s\n", user.Name)
}
