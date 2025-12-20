package day14

import (
	"aoc2025/internal/must"
	"math"
	"strings"
)

type Day14 struct {
	contestants []reindeer
}

func NewDay14(data []string) *Day14 {
	var r []reindeer
	for _, line := range data {
		r = append(r, *reindeerFrom(line))
	}
	return &Day14{contestants: r}
}

type reindeer struct {
	speed     int
	endurance int
	rest      int
}

func reindeerFrom(line string) *reindeer {
	fields := strings.Fields(line)
	return &reindeer{
		speed:     must.ParseInt(fields[3]),
		endurance: must.ParseInt(fields[6]),
		rest:      must.ParseInt(fields[13]),
	}
}

func (d *Day14) Part1() (int, error) {
	return d.Part1WithTime(2503)
}

func (d *Day14) Part1WithTime(raceTime int) (int, error) {
	best := 0

	for _, r := range d.contestants {
		iterations := raceTime / (r.endurance + r.rest)
		leftOverSeconds := raceTime % (r.endurance + r.rest)

		distance := iterations * r.endurance * r.speed
		distance += min(leftOverSeconds, r.endurance) * r.speed

		best = max(best, distance)
	}

	return best, nil
}

func (d *Day14) Part2() (int, error) {
	return d.Part2WithTime(2503)
}

func (d *Day14) Part2WithTime(raceTime int) (int, error) {
	points := make([]int, len(d.contestants))
	for t := 1; t <= raceTime; t++ {
		positions := make([]int, len(d.contestants))
		firstPlace := math.MinInt
		for i, r := range d.contestants {
			iterations := t / (r.endurance + r.rest)
			leftOverSeconds := t % (r.endurance + r.rest)

			distance := iterations * r.endurance * r.speed
			distance += min(leftOverSeconds, r.endurance) * r.speed
			firstPlace = max(firstPlace, distance)
			positions[i] = distance
		}
		for i, p := range positions {
			if p == firstPlace {
				points[i]++
			}
		}
	}

	winner := math.MinInt
	for _, p := range points {
		winner = max(winner, p)
	}
	return winner, nil
}
