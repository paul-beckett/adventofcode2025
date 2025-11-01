package cmd

import (
	"aoc2025/cmd/aoc2025"
	"fmt"

	"github.com/spf13/cobra"
)

var subCommands = []*cobra.Command{
	aoc2025.NewCommand(),
}

func NewRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:     "aoc",
		Example: "go run main.go aoc2025 day01",
		Run: func(cmd *cobra.Command, args []string) {
			for _, subCommand := range subCommands {
				fmt.Printf("running %s:\n", subCommand.Name())
				subCommand.Run(cmd, args)
			}
		},
	}
	for _, subCommand := range subCommands {
		root.AddCommand(subCommand)
	}

	return root
}
