package main

import (
	"log"
	"os"

	"github.com/dahakeadi15/gator/internal/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	programState := &state{
		cfg: &cfg,
	}
	cmds := commands{
		registeredCommands: make(map[string]func(s *state, cmd command) error),
	}

	cmds.register("login", handlerLogin)

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
