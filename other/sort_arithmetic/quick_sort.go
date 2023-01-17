package sort_arithmetic

func QuickSort(arr []int) []int {
	return _QuickSort(arr, 0, len(arr)-1)
}

func _QuickSort(arr []int, left, right int) []int {
	if right > left {
		partitionIndex := PartitionWay(arr, left, right)
		_QuickSort(arr, left, partitionIndex-1)
		_QuickSort(arr, partitionIndex+1, right)
	}
	return arr
}

// PartitionWay 单路快排-从左到右遍历
func PartitionWay(arr []int, left, right int) int {
	// 先分区，然后把基准放到边界上
	partition := left
	index := partition + 1
	for i := index; i <= right; i++ {
		if arr[partition] > arr[i] {
			// 当前值小于基准值就换
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}
	}
	arr[partition], arr[index-1] = arr[index-1], arr[partition]
	return index - 1
}

// Partition1Way 双路快排：双指针从首尾向中间移动
func Partition1Way(arr []int, low, height int) int {
	temp := arr[low] // 基准
	for low < height {
		// 当队尾的元素大于等于基准数据时，向前挪动height
		for low < height && arr[height] >= temp {
			height--
		}
		// 如果队尾元素小于基准元素，那么将其赋值给low
		arr[low] = arr[height]

		// 当队首的元素小于等于基准时，向后挪动low
		for low < height && arr[low] <= temp {
			low++
		}

		// 当队首元素大于基准时，将其赋值给height
		arr[height] = arr[low]
	}
	// 跳出循环时，height和low相等,此时的low活着height就是temp正确的索引位置
	// 由原理部分可以清楚的了解到low位置的值并不是temp,所以要将temp赋值给arr[low]
	arr[low] = temp
	return low
}
