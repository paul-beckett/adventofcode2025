package day00

import (
	"aoc2025/challenge/aoc2025/day00"
	"aoc2025/internal/file"
	"aoc2025/internal/metrics"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "day00",
		Run: func(cmd *cobra.Command, args []string) {
			day := day00.NewDay00(file.ReadFile("./input/aoc2025/day00.txt"))
			metrics.PrintResultAndTime("part1", day.Part1)
			metrics.PrintResultAndTime("part2", day.Part2)
		},
	}
}
