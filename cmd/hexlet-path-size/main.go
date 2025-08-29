package main

import (
	"context"
	"os"
)

func main() {
	cmd := GetAvailableCommands()

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		return
	}
}
