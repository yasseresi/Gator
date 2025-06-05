package main

import (
	"fmt"

	"github.com/yasseresi/Gator/internal/config"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("you have to add the username")
	}
	s.config.UserName = cmd.args[0]
	err := config.Write(*s.config)
	if err != nil {
		return fmt.Errorf("can't write the new username")
	}
	fmt.Println("your current username is : " + s.config.UserName)
	return nil
}
