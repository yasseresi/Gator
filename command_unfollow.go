package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yasseresi/Gator/internal/database"
)

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("unfollow command requires an argument: url")
	}

	params := database.DeleteFeedFollowForUserAndFeedUrlParams{
		UserID: user.ID,
		Url:    sql.NullString{String: cmd.args[0], Valid: true},
	}

	err := s.db.DeleteFeedFollowForUserAndFeedUrl(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to unfollow feed with URL '%s': %v", params.Url.String, err)
	}

	fmt.Printf("Successfully unfollowed feed: %s\n", params.Url.String)
	return nil
}
