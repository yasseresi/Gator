package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *state) error {
	fmt.Println("Starting feed scraping...")
	// Get the next feed to fetch from the DB
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get next feed to fetch: %v", err)
	}
	fmt.Printf("Found feed to fetch: %s\n", feed.Name.String)

	// Mark it as fetched
	fmt.Printf("Marking feed as fetched: %s\n", feed.ID)
	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %v", err)
	}
	fmt.Println("Feed marked as fetched successfully")

	// Fetch the feed using the URL
	fmt.Printf("Fetching RSS from URL: %s\n", feed.Url.String)
	rssFeed, err := fetchFeed(context.Background(), feed.Url.String)
	if err != nil {
		return fmt.Errorf("failed to fetch feed from URL '%s': %v", feed.Url.String, err)
	}

	// Iterate over the items in the feed and print their titles
	fmt.Printf("Fetched feed: %s\n", rssFeed.Channel.Title)
	fmt.Printf("Found %d items:\n", len(rssFeed.Channel.Item))

	for i, item := range rssFeed.Channel.Item {
		fmt.Printf("  %d. %s\n", i+1, item.Title)
	}

	return nil
}
