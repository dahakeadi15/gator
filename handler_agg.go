package main

import (
	"context"
	"fmt"
)

func handlerAggregator(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("feed: \n%+v\n", *feed)
	return nil
}
