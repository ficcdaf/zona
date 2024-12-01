package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ficcdaf/zona/internal/builder"
	"github.com/urfave/cli/v2"
)

func main() {
	var force bool
	var clean bool
	app := &cli.App{
		Name:    "zona",
		Version: "0.0.1",
		Usage:   "Static site builder.",
		Flags: []cli.Flag{
			cli.VersionFlag,
		},
		Commands: []*cli.Command{
			{
				Name:    "print",
				Usage:   "Prints helpful information.",
				Aliases: []string{"p"},
				Action: func(ctx *cli.Context) error {
					fmt.Println("Printing some good stuff.")
					return nil
				},
			},
			{
				Name:      "build",
				Usage:     "Builds the website.",
				HideHelp:  true,
				Args:      true,
				UsageText: "zona build [opts] input_dir (output_dir)",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "force",
						Aliases:     []string{"f"},
						Usage:       "Overwrites files in output_dir if they exist.",
						Destination: &force,
					},
					&cli.BoolFlag{
						Name:        "clean",
						Aliases:     []string{"c"},
						Usage:       "Remove all files in output_dir that are not results of the current build.",
						Destination: &clean,
					},
				},
				Description: `The build command takes an input directory and an optional output directory. The default output_dir is named "{input_name}-built" and written to the same parent directory as the input.`,
				Aliases:     []string{"b"},
				Action: func(ctx *cli.Context) error {
					if ctx.NArg() == 0 {
						cli.ShowCommandHelpAndExit(ctx, "build", 1)
					}
					if force {
						fmt.Println("I'll need to pass the force var as an argument to build...")
					}
					return nil
					// err := build(ctx.Args().Get(0), ctx.Args().Get(1))
					// return err
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func build(rootPath string, outPath string, force bool, clean bool) error {
	settings := builder.GetSettings(rootPath, outPath, force, clean)
	err := builder.Traverse(rootPath, outPath, settings)
	if err != nil {
		return fmt.Errorf("Build error: %s", err)
	}
	return nil
}
