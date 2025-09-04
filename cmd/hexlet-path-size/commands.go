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
				Usage:   "For get humanrized value",
			},
			&cli.BoolFlag{
				Name:    "all, a",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "include hidden files and directories.",
			},
		},
		Action: func(ctx context.Context, command *cli.Command) error {
			h := command.Bool("human")
			a := command.Bool("all")
			p := command.Args().Get(0)

			if p == "" {
				return fmt.Errorf("path is required")
			}

			size, err := GetSize(p, h, a)
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
