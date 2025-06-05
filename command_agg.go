package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	// Feed URL to fetch
	feedURL := "https://www.wagslane.dev/index.xml"

	// Fetch the RSS feed
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %v", err)
	}

	// Print the entire struct to console
	fmt.Printf("Feed: %+v\n", *feed)

	return nil
}
