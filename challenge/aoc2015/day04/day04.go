package day04

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

type Day04 struct {
	data []string
}

func NewDay04(data []string) *Day04 {
	return &Day04{data: data}
}

func (d *Day04) Part1() (int, error) {
	return findHash(d.data[0], 1, 5), nil
}

func findHash(prefix string, i, zeros int) int {
	h := md5.New()
	target := strings.Repeat("0", zeros)
	for {
		h.Reset()
		io.WriteString(h, fmt.Sprintf("%s%d", prefix, i))
		res := fmt.Sprintf("%x", h.Sum(nil))
		if strings.HasPrefix(res, target) {
			return i
		}
		i++
	}
}

func (d *Day04) Part2() (int, error) {
	return findHash(d.data[0], 1, 6), nil
}
