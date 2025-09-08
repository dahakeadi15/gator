package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dahakeadi15/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) > 1 {
		return fmt.Errorf("usage: %s <limit(optional)>", cmd.Name)
	}

	limit := 2
	if len(cmd.Arguments) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.Arguments[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}
		limit = parsedLimit
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user %s: %w", user.Name, err)
	}

	if len(posts) == 0 {
		fmt.Printf("User %s has no posts.", user.Name)
		return nil
	}

	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 1"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("===========================")
	}

	return nil
}
