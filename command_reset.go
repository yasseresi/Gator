package main

import (
	"context"
	"fmt"
	"os"
)

func handlReset(s *state, cmd command) error {

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
	fmt.Println("Users deleted successfuly")
	return nil
}
