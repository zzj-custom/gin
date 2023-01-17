package sort_arithmetic

import "github.com/gookit/goutil/dump"

// SelectionSort 选择排序
// message:
// 1. 初始序列arr，无序。分成有序区和无序区，有序区初始为0，不断变大；无序区初始为len(arr)，不断变小。
// 2. 遍历无序找到最小值，与无序区最左边交换。有序区长度+1。
// 3. 重复第二步
func SelectionSort(list []int) []int {
	length := len(list)
	for i := 0; i < length-1; i++ {
		arrIndex := i
		for j := i + 1; j < length; j++ {
			if list[j] < list[arrIndex] {
				arrIndex = j
			}
		}
		list[i], list[arrIndex] = list[arrIndex], list[i]
		dump.P(list)
	}
	return list
}
