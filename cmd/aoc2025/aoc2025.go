package aoc2025

import (
	"aoc2025/cmd/aoc2025/day01"
	"aoc2025/cmd/aoc2025/day02"
	"aoc2025/cmd/aoc2025/day03"
	"aoc2025/cmd/aoc2025/day04"
	"aoc2025/cmd/aoc2025/day05"
	"aoc2025/cmd/aoc2025/day06"
	"aoc2025/cmd/aoc2025/day07"
	"aoc2025/cmd/aoc2025/day08"
	"aoc2025/cmd/aoc2025/day09"
	"aoc2025/cmd/aoc2025/day10"
	"aoc2025/cmd/aoc2025/day11"
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
	day10.NewCommand(),
	day11.NewCommand(),
}

func NewCommand() *cobra.Command {
	year := &cobra.Command{
		Use: "2025",
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
