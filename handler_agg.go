package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dahakeadi15/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAggregator(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>\nwhere, time_between_reqs = 1s, 1m, 1h, etc", cmd.Name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("couldn't parse duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s...\n", timeBetweenReqs.String())

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("couldn't get next feed to fetch: %w", err)
		return
	}

	log.Println("Found a feed to fetch!")
	scrapeFeed(s.db, nextFeed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("couldn't mark feed as fetched: %v", err)
		return
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	fmt.Printf("Fetched %d titles from feed %s:\n", len(rssFeed.Channel.Item), rssFeed.Channel.Title)
	count := 0
	for _, post := range rssFeed.Channel.Item {
		publishedAt := sql.NullTime{}
		if parsedTime, err := time.Parse(time.RFC1123Z, post.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  parsedTime,
				Valid: true,
			}
		}

		newPost, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     post.Title,
			Url:       post.Link,
			Description: sql.NullString{
				String: post.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key value violates unique constraint \"posts_url_key\"") {
				log.Printf("couldn't add post (%s): %v\n", post.Title, err)
			}
			continue
		}

		log.Printf("added post: %s\n", newPost.Title)
		count++
	}
	fmt.Printf("Added %d new posts from %s.\n", count, feed.Name)
	fmt.Println("==============================")
}
