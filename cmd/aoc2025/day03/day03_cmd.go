package day03

import (
	"aoc2025/challenge/aoc2025/day03"
	"aoc2025/internal/file"
	"aoc2025/internal/metrics"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "day03",
		Run: func(cmd *cobra.Command, args []string) {
			day := day03.NewDay03(file.ReadFile("./input/aoc2025/day03.txt"))
			metrics.PrintResultAndTime("part1", day.Part1)
			metrics.PrintResultAndTime("part2", day.Part2)
		},
	}
}
