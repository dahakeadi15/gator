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

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed Follow created:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("=================================")

	return nil
}

func handlerListFeedFollows(s *state, cmd command) error {
	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currUser.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feeds the current user follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feeds followed by user %s\n:", currUser.Name)
	for _, ff := range feedFollows {
		fmt.Printf(" * %s\n", ff.FeedName)
	}
	fmt.Println("===============================")

	return nil
}

func printFeedFollow(userName, feedName string) {
	fmt.Printf(" * UserName:  %s\n", userName)
	fmt.Printf(" * FeedName:  %s\n", feedName)
}
