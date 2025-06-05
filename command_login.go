package main

import (
	"context"
	"fmt"
	"os"

	"github.com/yasseresi/Gator/internal/config"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("you have to add the username")
	}
	username := cmd.args[0]

	// Try to get the user from the database
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		fmt.Printf("user '%s' not found", username)
		os.Exit(1)
	}

	// Set the current user in the config
	s.config.UserName = user.Name
	err = config.Write(*s.config)
	if err != nil {
		return fmt.Errorf("failed to update config: %v", err)
	}
	fmt.Printf("User set to: %s\n", s.config.UserName)
	return nil
}
