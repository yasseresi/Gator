package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/yasseresi/Gator/internal/config"
	"github.com/yasseresi/Gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("register command requires a username")
	}

	name := cmd.args[0]

	// Create user parameters
	createUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	// Create the user in the database
	user, err := s.db.CreateUser(context.Background(), createUserParams)
	if err != nil {
		// Check if this is a unique constraint violation (user already exists)
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_name_key\"" {
			fmt.Fprintf(os.Stderr, "User '%s' already exists\n", name)
			os.Exit(1)
		}
		return fmt.Errorf("failed to create user: %v", err)
	}

	// Set the current user in the config
	s.config.UserName = name
	err = config.Write(*s.config)
	if err != nil {
		return fmt.Errorf("failed to update config: %v", err)
	}

	// Print success message and debug info
	fmt.Printf("User '%s' created successfully!\n", name)
	fmt.Printf("User data: ID=%s, Name=%s, CreatedAt=%s\n",
		user.ID, user.Name, user.CreatedAt.Format(time.RFC3339))

	return nil
}
