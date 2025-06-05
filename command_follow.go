package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yasseresi/Gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("follow command requires a URL argument")
	}

	url := cmd.args[0]

	// Get the current user from the database
	user, err := s.db.GetUser(context.Background(), s.config.UserName)
	if err != nil {
		return fmt.Errorf("user '%s' not found: %v", s.config.UserName, err)
	}

	// Look up the feed by URL
	feed, err := s.db.GetFeedByURL(context.Background(), sql.NullString{String: url, Valid: true})
	if err != nil {
		return fmt.Errorf("feed with URL '%s' not found: %v", url, err)
	}

	// Create the feed follow parameters
	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	// Create the feed follow record
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("failed to follow feed: %v", err)
	}

	// Print success message with feed name and user name
	fmt.Printf("User '%s' is now following feed '%s'\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
