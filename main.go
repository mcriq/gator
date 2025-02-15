package main

import (
	"log"
	"os"

	"github.com/mcriq/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read config file: %v\n", err)
		os.Exit(1)
	}

	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command)error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}