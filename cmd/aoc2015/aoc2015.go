package aoc2015

import (
	"aoc2025/cmd/aoc2015/day01"
	"aoc2025/cmd/aoc2015/day02"
	"aoc2025/cmd/aoc2015/day03"
	"fmt"

	"github.com/spf13/cobra"
)

var subCommands = []*cobra.Command{
	day01.NewCommand(),
	day02.NewCommand(),
	day03.NewCommand(),
}

func NewCommand() *cobra.Command {
	year := &cobra.Command{
		Use: "2015",
		Run: func(cmd *cobra.Command, args []string) {
			for _, subCommand := range subCommands {
				fmt.Printf("running %s:\n", subCommand.Name())
				subCommand.Run(cmd, args)
			}
		},
	}
	for _, cmd := range subCommands {
		year.AddCommand(cmd)
	}
	return year
}
