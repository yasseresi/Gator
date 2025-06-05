package main

import (
	"fmt"
	"os"

	"github.com/yasseresi/Gator/internal/config"
)

func main() {

	conf := config.ReadConfig()

	s := state{
		config: &conf,
	}
	cmds := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "You have to enter more arguments to execute this command")
		os.Exit(1)
	}
	command_name := args[1]
	arguments := args[2:]

	cmd := command{
		name: command_name,
		args: arguments,
	}
	if err := cmds.run(&s, cmd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
