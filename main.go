package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/dahakeadi15/gator/internal/config"
	"github.com/dahakeadi15/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(s *state, cmd command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAggregator)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFeedFollow))
	cmds.register("unfollow", middlewareLoggedIn(handlerFeedUnfollow))
	cmds.register("following", middlewareLoggedIn(handlerListFeedFollows))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	cliArgs := os.Args
	if len(cliArgs) < 2 {
		log.Fatal("Usage: cli <command> [...args]")
	}

	cmdName := cliArgs[1]
	cmdArgs := cliArgs[2:]

	cmd := command{
		Name:      cmdName,
		Arguments: cmdArgs,
	}

	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
