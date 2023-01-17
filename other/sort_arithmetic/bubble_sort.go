package sort_arithmetic

// BubbleSort 冒泡排序
// message: 通过比较前后两个元素的大小，然后交换位置来实现排序。
// 每次比较相邻两个数的大小，如果前面的数大于后面的数，则交换两个数的位置（否则不变），向后移动
func BubbleSort(list []int) []int {
	length := len(list)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}

// BubbleSort1 冒泡排序优化
// message:
func BubbleSort1(list []int) []int {
	length := len(list)
	for i := length - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if list[j] < list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}
