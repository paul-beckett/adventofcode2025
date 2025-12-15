package aoc2015

import (
	"aoc2025/cmd/aoc2015/day01"
	"aoc2025/cmd/aoc2015/day02"
	"aoc2025/cmd/aoc2015/day03"
	"aoc2025/cmd/aoc2015/day04"
	"aoc2025/cmd/aoc2015/day05"
	"aoc2025/cmd/aoc2015/day06"
	"aoc2025/cmd/aoc2015/day07"
	"aoc2025/cmd/aoc2015/day08"
	"aoc2025/cmd/aoc2015/day09"
	"fmt"

	"github.com/spf13/cobra"
)

var subCommands = []*cobra.Command{
	day01.NewCommand(),
	day02.NewCommand(),
	day03.NewCommand(),
	day04.NewCommand(),
	day05.NewCommand(),
	day06.NewCommand(),
	day07.NewCommand(),
	day08.NewCommand(),
	day09.NewCommand(),
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
