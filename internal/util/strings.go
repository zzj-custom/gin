package util

import (
	"github.com/gin-gonic/gin"
	"math"
)

func GetIP(ctx *gin.Context) string {
	forwarded := ctx.Request.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return ctx.Request.RemoteAddr
}

func Round(x float64, unit int) float64 {
	p := math.Pow10(unit)
	return float64(int64(x*p+0.5)) / p
}

func BinarySearch(arr []int, search int, left, right int) bool {
	if left > right {
		return false
	}
	middle := (left + right) / 2

	if arr[middle] == search {
		return true
	} else if arr[middle] > search {
		return BinarySearch(arr, search, 0, middle-1)
	} else {
		return BinarySearch(arr, search, middle+1, right)
	}
}

func twoSum(arr []int, target int) []int {
	m := make(map[int]int)
	for k, val := range arr {
		another := target - val
		if _, ok := m[another]; ok {
			return []int{k, m[another]}
		}
		m[val] = k
	}
	return nil
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++

		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
