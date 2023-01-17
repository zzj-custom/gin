package util

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestRound(t *testing.T) {
	var (
		a = 3.14
		b = 3.1415
		c = 3.1415926
		d = -3.14
		e = 1.035
	)

	// go语言四舍五入最大的问题是有时候5不进位
	fmt.Println(Round(a, 1))
	fmt.Println(Round(b, 2))
	fmt.Println(Round(c, 4))
	fmt.Println(Round(d, 1))
	fmt.Println(Round(e, 2))
}

func TestQuickSort(t *testing.T) {
	s := make([]int, 0, 10000)
	for i := 1; i <= 10000; i++ {
		s = append(s, i)
	}
	start := time.Now().UnixNano()
	BinarySearch(s, 3875, 0, len(s)-1)
	fmt.Println(time.Now().UnixNano() - start)
}

func TestTwoSum(t *testing.T) {
	s := []int{2, 7, 9, 11}
	target := 9
	assert.Equal(t, twoSum(s, target), []int{1, 0})
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"lengthOfLongestSubstring",
			args{s: "pwwkew"},
			3,
		},
		{
			name: "lengthOfLongestSubstring",
			args: args{s: "acdrfaedc"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
