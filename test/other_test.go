package test

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[int]int, 10)
	for i := 1; i <= 10; i++ {
		m[i] = i
	}

	// 如果不传参会导致只能获取到一个相同的
	for k, v := range m {
		go func(k, v int) {
			fmt.Println("k ->", k, "v ->", v)
		}(k, v)
	}
}

func TestArray(t *testing.T) {
	d := [...]int{1, 2, 4: 8, 10}
	c := [...]int{1, 2, 3, 4: 5, 6}
	dump.P(d, c, len(d))
	for v := range c {
		dump.P(v)
	}
}

func TestSlice(t *testing.T) {
	i := 3
	s := []int{1, 2, 3, 5, 6, 7}
	s = append(s, 0)
	dump.P(s)
	copy(s[i+1:], s[i:])
	s[i] = 4
	dump.P(s)
}
