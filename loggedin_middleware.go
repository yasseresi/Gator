package main

import (
	"context"
	"fmt"

	"github.com/yasseresi/Gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		// Get username from config
		username := s.config.UserName

		// Try to get the user from the database
		user, err := s.db.GetUser(context.Background(), username)
		if err != nil {
			return fmt.Errorf("user '%s' not found: %v", username, err)
		}

		// Call the original handler with the user
		return handler(s, cmd, user)
	}
}
