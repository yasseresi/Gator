package main

import (
	"context"
	"fmt"

	"github.com/yasseresi/Gator/internal/database"
)

func handleFollowing(s *state, cmd command, user database.User) error {
	username := user.Name

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
