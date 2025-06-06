package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yasseresi/Gator/internal/database"
)

func handleAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("addfeed command requires two arguments: name and url")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	// Create the feed parameters with correct nullable types
	new_feed := database.CreateFeddParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      sql.NullString{String: name, Valid: true},
		Url:       sql.NullString{String: url, Valid: true},
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
	}

	// Create the feed in the database
	feed, err := s.db.CreateFedd(context.Background(), new_feed)
	if err != nil {
		return fmt.Errorf("failed to create feed '%s': %v", name, err)
	}

	// Create feed follow automatically
	new_feed_follow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feed_follow, err := s.db.CreateFeedFollow(context.Background(), new_feed_follow)
	if err != nil {
		return fmt.Errorf("failed to create feed follow '%s': %v", name, err)
	}

	// Print the feed details
	fmt.Printf("Feed created successfully!\n")
	fmt.Printf("User '%s' is now automatically following the feed.\n", feed_follow.UserName)
	fmt.Printf("ID: %s\n", feed.ID)
	fmt.Printf("Name: %s\n", feed.Name.String)
	fmt.Printf("URL: %s\n", feed.Url.String)
	fmt.Printf("User ID: %s\n", feed.UserID.UUID)
	fmt.Printf("Created At: %s\n", feed.CreatedAt.Format(time.RFC3339))

	return nil
}
