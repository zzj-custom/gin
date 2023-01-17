package sort_arithmetic

import (
	"math"
)

// BucketSort 桶排序
func BucketSort(arr []int, bucketSize int) []int {
	// 获取最值
	length := len(arr)
	minNum, _, maxValue := getMost(arr)

	// 初始化桶
	bucketNum := maxValue/bucketSize + 1
	buckets := make([][]int, bucketNum)
	for i := 0; i < bucketNum; i++ {
		buckets[i] = make([]int, 0)
	}

	// 将元素映射分配到每个桶里面
	for i := 0; i < length; i++ {
		bucketId := (arr[i] - minNum) / bucketSize
		buckets[bucketId] = append(buckets[bucketId], arr[i])
	}
	// 对每个桶进行排序
	arrIndex := 0
	for i := 0; i < bucketNum; i++ {
		if len(buckets[i]) == 0 {
			// 空桶
			continue
		}
		InsertionSort(buckets[i])
		for j := 0; j < len(buckets[i]); j++ {
			arr[arrIndex] = buckets[i][j]
			arrIndex++
		}
	}
	return arr
}

func getMost(arr []int) (int, int, int) {
	minValue, maxValue := math.MaxInt64, math.MinInt64
	for i := 0; i < len(arr); i++ {
		if minValue > arr[i] {
			minValue = arr[i]
		}
		if maxValue < arr[i] {
			maxValue = arr[i]
		}
	}
	return minValue, maxValue, maxValue - minValue
}
