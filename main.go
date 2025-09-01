package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dahakeadi15/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	s := state{
		cfg: &cfg,
	}
	cmds := commands{
		functions: map[string]func(s *state, cmd command) error{},
	}

	cmds.register("login", handlerLogin)

	cliArgs := os.Args
	if len(cliArgs) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	cmdName := cliArgs[1]
	cmdArgs := cliArgs[2:]

	cmd := command{
		name:      cmdName,
		arguments: cmdArgs,
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
