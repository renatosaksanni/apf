package main

import (
	"log"
	"os"

	"github.com/renatosaksanni/apf/config"
	"github.com/renatosaksanni/apf/internal/interfaces/cli"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	if err := cli.Run(os.Args); err != nil {
		log.Fatalf("could not run command: %v", err)
	}
}
