package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dahakeadi15/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFeedFollow(s *state, cmd command) error {
	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find current user: %w", err)
	}

	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}
	feedUrl := cmd.Arguments[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	newFeedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed Follow created successfully:")
	fmt.Printf(" * UserName:  %s\n", newFeedFollow.UserName)
	fmt.Printf(" * FeedName:  %s\n", newFeedFollow.FeedName)
	fmt.Println("=================================")

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get feeds the current user is following: %w", err)
	}

	fmt.Println("Feeds followed by current user:")
	for _, feed := range following {
		fmt.Printf(" * %s\n", feed.FeedName)
	}
	fmt.Println("===============================")

	return nil
}
