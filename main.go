package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/yasseresi/Gator/internal/config"
	"github.com/yasseresi/Gator/internal/database"
)

func main() {

	conf := config.ReadConfig()

	// Use the database URL from config, not hostname
	db, err := sql.Open("postgres", conf.DbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %v\n", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	s := state{
		config: &conf,
		db:     dbQueries,
	}

	cmds := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlReset)
	cmds.register("users", handleList)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handleAddFeed))
	cmds.register("feeds", handleFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handleFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handleUnfollow))
	cmds.register("browse", middlewareLoggedIn(handleBrowse))

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
