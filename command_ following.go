package main

import (
	"context"
	"fmt"
)

func handleFollowing(s *state, cmd command) error {
	username := s.config.UserName

	// Try to get the user from the database
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user '%s' not found: %v", username, err)
	}

	// Get all feeds followed by the user
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get feeds for user '%s': %v", username, err)
	}

	if len(feeds) == 0 {
		fmt.Printf("User '%s' is not following any feeds.\n", username)
		return nil
	}

	fmt.Printf("Feeds followed by '%s':\n", username)
	for _, feed := range feeds {
		fmt.Printf("- %s (%s)\n", feed.FeedName.String, feed.FeedUrl.String)
	}

	return nil
}
