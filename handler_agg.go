package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dahakeadi15/gator/internal/database"
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
	for _, rssItem := range rssFeed.Channel.Item {
		fmt.Println(rssItem.Title)
	}
	fmt.Println("==============================")
}
