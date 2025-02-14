package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mcriq/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read config file: %v\n", err)
		os.Exit(1)
	}

	err = cfg.SetUser("Matt")
	if err != nil {
		log.Fatalf("unable to set user: %v\n", err)
		os.Exit(1)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("unable to read config file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(cfg.DBURL)
}