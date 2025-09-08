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
			return fmt.Errorf("limit must be an integer: %w", err)
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

	for i, post := range posts {
		fmt.Printf("%d:\n", i+1)
		printPost(post)
	}

	return nil
}

func printPost(post database.Post) {
	fmt.Printf(" * Title:        %s\n", post.Title)
	fmt.Printf(" * URL:          %s\n", post.Url)
	fmt.Printf(" * Description:  %s\n", post.Description.String)
	fmt.Printf(" * Published At: %s\n", post.PublishedAt)
}
