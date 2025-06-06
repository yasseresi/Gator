package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yasseresi/Gator/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2) // default limit

	// Parse limit if provided
	if len(cmd.args) >= 1 {
		parsedLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid limit '%s': %v", cmd.args[0], err)
		}
		if parsedLimit <= 0 {
			return fmt.Errorf("limit must be a positive number")
		}
		limit = int32(parsedLimit)
	}

	// Get posts for the user
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("failed to get posts: %v", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found. Follow some feeds and run the aggregator first!")
		return nil
	}

	fmt.Printf("Found %d posts for user %s:\n\n", len(posts), user.Name)

	for i, post := range posts {
		fmt.Printf("* %s\n", post.Title)
		fmt.Printf("  URL: %s\n", post.Url)

		if post.Description.Valid && post.Description.String != "" {
			// Truncate description if it's too long
			description := post.Description.String
			if len(description) > 200 {
				description = description[:200] + "..."
			}
			fmt.Printf("  Description: %s\n", description)
		}

		if post.PublishedAt.Valid {
			fmt.Printf("  Published: %s\n", post.PublishedAt.Time.Format("2006-01-02 15:04:05"))
		}

		if post.FeedName.Valid {
			fmt.Printf("  Feed: %s\n", post.FeedName.String)
		}

		// Add separator between posts except for the last one
		if i < len(posts)-1 {
			fmt.Println()
		}
	}

	return nil
}
