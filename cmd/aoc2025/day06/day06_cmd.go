package day06

import (
	"aoc2025/challenge/aoc2025/day06"
	"aoc2025/internal/file"
	"aoc2025/internal/metrics"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "day06",
		Run: func(cmd *cobra.Command, args []string) {
			day := day06.NewDay06(file.ReadFile("./input/aoc2025/day06.txt"))
			metrics.PrintResultAndTime("part1", day.Part1)
			metrics.PrintResultAndTime("part2", day.Part2)
		},
	}
}
