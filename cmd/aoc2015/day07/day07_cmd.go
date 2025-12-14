package day07

import (
	"aoc2025/challenge/aoc2015/day07"
	"aoc2025/internal/file"
	"aoc2025/internal/metrics"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "day07",
		Run: func(cmd *cobra.Command, args []string) {
			day := day07.NewDay07(file.ReadFile("./input/aoc2015/day07.txt"))
			metrics.PrintResultAndTime("part1", day.Part1)
			metrics.PrintResultAndTime("part2", day.Part2)
		},
	}
}
