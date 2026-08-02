package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var version = "unknown"

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "print the version",
	}

	cmd := &cli.Command{
		Name:    "template",
		Version: version,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
