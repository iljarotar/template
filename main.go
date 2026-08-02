package main

import (
	"context"
	"log"
	"os"

	"github.com/iljarotar/template/renderer"
	"github.com/urfave/cli/v3"
)

var version = "unknown"

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print the version",
	}

	cmd := &cli.Command{
		Name:    "template",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "template",
				DefaultText: "path to a template file",
				Usage:       "specify path to a template file",
				Required:    true,
				Value:       "",
				Aliases:     []string{"t"},
			},
			&cli.StringFlag{
				Name:        "output",
				DefaultText: "path to the output file",
				Usage:       "specify path to the output file",
				Required:    true,
				Value:       "",
				Aliases:     []string{"o"},
			},
			&cli.StringFlag{
				Name:        "data",
				DefaultText: "path to a data file in json format",
				Usage:       "specify path to a data file",
				Value:       "",
				Aliases:     []string{"d"},
			},
		},
		Usage:       "renders go templates",
		UsageText:   "template -t <template-file> -o <output-file> (-d <data-file>)",
		Description: "template renderer for go templates",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			tmpl := cmd.String("template")
			data := cmd.String("data")
			out := cmd.String("output")

			return renderer.Render(tmpl, data, out)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
