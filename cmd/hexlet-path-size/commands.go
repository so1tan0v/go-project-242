package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func GetAvailableCommands() cli.Command {
	return cli.Command{
		Name:    "hexlet-path-size",
		Version: "0.0.1",
		Usage:   "./hexlet-path-size --human [some file name]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human, H",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit) (default: false)",
			},
			&cli.BoolFlag{
				Name:    "all, a",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "include hidden files and directories (default: false)",
			},
			&cli.BoolFlag{
				Name:    "recursive, r",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "recursive size of directories (default: false)",
			},
		},
		Action: func(ctx context.Context, command *cli.Command) error {
			pathToObject := command.Args().Get(0)
			human := command.Bool("human")
			all := command.Bool("all")
			recursive := command.Bool("recursive")

			if pathToObject == "" {
				return fmt.Errorf("path is required")
			}

			size, err := GetResult(pathToObject, human, all, recursive)
			if err != nil {
				return fmt.Errorf("error getting size: %w", err)
			}

			fmt.Println(size)

			return nil
		},
	}
}

//func getCommandList() {
//
//}
