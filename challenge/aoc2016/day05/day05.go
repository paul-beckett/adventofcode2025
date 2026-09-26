package day05

import (
	"crypto/md5"
	"fmt"
	"io"
	"slices"
	"strings"
)

type Day05 struct {
	data []string
}

func NewDay05(data []string) *Day05 {
	return &Day05{data: data}
}

func (d *Day05) Part1() (string, error) {
	i := 0
	door := d.data[0]
	target := "00000"
	var password []rune
	for len(password) < 8 {
		var s string
		i, s = findNextHash(target, door, i)
		password = append(password, rune(s[5]))
		i++
	}
	return string(password), nil
}

func findNextHash(target string, door string, i int) (int, string) {
	h := md5.New()
	for {
		h.Reset()
		io.WriteString(h, fmt.Sprintf("%s%d", door, i))
		s := fmt.Sprintf("%x", h.Sum(nil))
		if strings.HasPrefix(s, target) {
			return i, s
		}
		i++
	}
}

func (d *Day05) Part2() (string, error) {
	password := make([]rune, 8)

	i := 0
	door := d.data[0]
	target := "00000"
	for {
		var s string
		i, s = findNextHash(target, door, i)
		pos := int(s[5] - '0')
		if pos < len(password) && password[pos] == 0 {
			password[pos] = rune(s[6])
		}
		i++
		if !slices.Contains(password, 0) {
			break
		}
	}
	//has to run about 25million hashes! Completes in 12.5s on Ryzen 3700X
	return string(password), nil
}
