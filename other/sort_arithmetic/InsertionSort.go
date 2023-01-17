package sort_arithmetic

// InsertionSort 插入排序
// message:
// 1. 将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列
// 2. 从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
// 3. 如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。
func InsertionSort(arr []int) []int {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
	return arr
}

func InsertionSort1(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		index := i - 1
		current := arr[i]
		for index >= 0 && arr[index] > current {
			arr[index+1] = arr[index]
			index--
		}
		arr[index] = current
	}
	return arr
}
