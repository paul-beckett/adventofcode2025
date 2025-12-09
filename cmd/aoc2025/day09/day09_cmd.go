package day09

import (
	"aoc2025/challenge/aoc2025/day09"
	"aoc2025/internal/file"
	"aoc2025/internal/metrics"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "day09",
		Run: func(cmd *cobra.Command, args []string) {
			day := day09.NewDay09(file.ReadFile("./input/aoc2025/day09.txt"))
			metrics.PrintResultAndTime("part1", day.Part1)
			metrics.PrintResultAndTime("part2", day.Part2)
		},
	}
}
