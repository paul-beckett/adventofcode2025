package day13

import (
	"aoc2025/challenge/aoc2015/day13"
	"aoc2025/internal/file"
	"aoc2025/internal/metrics"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "day13",
		Run: func(cmd *cobra.Command, args []string) {
			day := day13.NewDay13(file.ReadFile("./input/aoc2015/day13.txt"))
			metrics.PrintResultAndTime("part1", day.Part1)
			metrics.PrintResultAndTime("part2", day.Part2)
		},
	}
}
