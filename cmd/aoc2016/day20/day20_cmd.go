package day20

import (
	"aoc2025/challenge/aoc2016/day20"
	"aoc2025/internal/file"
	"aoc2025/internal/metrics"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "day20",
		Run: func(cmd *cobra.Command, args []string) {
			day := day20.NewDay20(file.ReadFile("./input/aoc2016/day20.txt"))
			metrics.PrintResultAndTime("part1", day.Part1)
			metrics.PrintResultAndTime("part2", day.Part2)
		},
	}
}
