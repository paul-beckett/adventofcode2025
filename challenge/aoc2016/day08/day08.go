package day08

import (
	"aoc2025/internal/math"
	"aoc2025/internal/must"
	"strings"
	"unicode"
)

type Day08 struct {
	data []string
}

func NewDay08(data []string) *Day08 {
	return &Day08{data: data}
}

const (
	width  = 50
	height = 6
)

func (d *Day08) drawScreen() map[math.Vector2]bool {
	screen := make(map[math.Vector2]bool)
	for _, instr := range d.data {
		fields := strings.FieldsFunc(instr, func(r rune) bool {
			return !unicode.IsDigit(r) && !unicode.IsLetter(r)
		})
		switch {
		case fields[0] == "rect":
			r := strings.Split(fields[1], "x")
			rect(screen, must.ParseInt(r[0]), must.ParseInt(r[1]))
		case fields[1] == "row":
			rotateRow(screen, must.ParseInt(fields[3]), must.ParseInt(fields[5]))
		case fields[1] == "column":
			rotateCol(screen, must.ParseInt(fields[3]), must.ParseInt(fields[5]))
		}
	}
	return screen
}

func (d *Day08) Part1() (int, error) {
	screen := d.drawScreen()
	return len(screen), nil
}

func (d *Day08) Part2() (string, error) {
	screen := d.drawScreen()

	var output []string
	for y := range height {
		row := make([]string, width)
		for x := range width {
			if screen[math.Vector2{X: x, Y: y}] {
				row[x] = "#"
			} else {
				row[x] = " "
			}
		}
		output = append(output, strings.Join(row, ""))
	}
	return "\n" + strings.Join(output, "\n"), nil
}

func rect(screen map[math.Vector2]bool, w, h int) {
	for x := range w {
		for y := range h {
			screen[math.Vector2{X: x, Y: y}] = true
		}
	}
}

func rotateRow(screen map[math.Vector2]bool, row, n int) {
	var pixels []math.Vector2
	for pixel := range screen {
		if pixel.Y == row {
			pixels = append(pixels, pixel)
		}
	}
	rotate(screen, pixels, math.Vector2{X: n, Y: 0})
}

func rotateCol(screen map[math.Vector2]bool, col, n int) {
	var pixels []math.Vector2
	for pixel := range screen {
		if pixel.X == col {
			pixels = append(pixels, pixel)
		}
	}
	rotate(screen, pixels, math.Vector2{X: 0, Y: n})
}

func rotate(screen map[math.Vector2]bool, items []math.Vector2, move math.Vector2) {
	var insert []math.Vector2
	for _, item := range items {
		delete(screen, item)
		next := item.Add(move)
		next.X %= width
		next.Y %= height
		insert = append(insert, next)
	}
	for _, item := range insert {
		screen[item] = true
	}
}
