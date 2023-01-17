package sort_arithmetic

import (
	"math"
)

// CountingSort 计数排序
// 1. 找出待排序的数组中最大和最小的元素
// 2. 统计数组中每个值为i的元素出现的次数，存入数组C的第i项
// 3. 对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）
// 4. 反向填充目标数组：将每个元素i放在新数组的第C(i)项，每放一个元素就将C(i)减去1
func CountingSort(arr []int) []int {
	arrLen := len(arr)
	maxValue := getMaxValue(arr)
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)
	for i := 0; i < arrLen; i++ {
		bucket[arr[i]] += 1
	}

	sortIndex := 0
	for j := 0; j < arrLen; j++ {
		for bucket[j] > 0 {
			arr[sortIndex] = j
			sortIndex++
			bucket[j]--
		}
	}
	return arr
}

func getMaxValue(arr []int) int {
	largest := math.MinInt64
	smallest := math.MaxInt64
	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}
	return largest - smallest
}
