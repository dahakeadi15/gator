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

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
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
	fmt.Printf(" * UserName:  %s\n", feed_follow.UserName)
	fmt.Printf(" * FeedName:  %s\n", feed_follow.FeedName)
	fmt.Println("=================================")

	return nil
}
