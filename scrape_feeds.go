package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/yasseresi/Gator/internal/database"
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

	// Iterate over the items in the feed and save them to the database
	fmt.Printf("Fetched feed: %s\n", rssFeed.Channel.Title)
	fmt.Printf("Found %d items:\n", len(rssFeed.Channel.Item))

	postsCreated := 0
	for _, item := range rssFeed.Channel.Item {
		// Parse the published date
		var publishedAt sql.NullTime
		if item.PubDate != "" {
			// Try common RSS date formats
			dateFormats := []string{
				time.RFC1123Z,
				time.RFC1123,
				"Mon, 2 Jan 2006 15:04:05 -0700",
				"Mon, 2 Jan 2006 15:04:05 MST",
				"2006-01-02T15:04:05Z07:00",
				"2006-01-02T15:04:05Z",
				"2006-01-02 15:04:05",
			}

			for _, format := range dateFormats {
				if parsedTime, err := time.Parse(format, strings.TrimSpace(item.PubDate)); err == nil {
					publishedAt = sql.NullTime{Time: parsedTime, Valid: true}
					break
				}
			}
		}

		// Create the post
		postParams := database.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		}

		_, err := s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			// If it's a unique constraint violation (post URL already exists), just skip it
			if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "UNIQUE constraint") {
				fmt.Printf("  Post already exists: %s\n", item.Title)
				continue
			}
			// For other errors, log them but continue processing
			log.Printf("Failed to create post '%s': %v", item.Title, err)
			continue
		}

		postsCreated++
		fmt.Printf("  Created post: %s\n", item.Title)
	}

	fmt.Printf("Successfully created %d new posts\n", postsCreated)
	return nil
}
