package sort_arithmetic

// MergeSort 归并排序
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[:middle]
	right := arr[middle:]
	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left, right []int) []int {
	res := make([]int, 0)
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	// 当任意一个left和right为空是将另外一个的剩余部分放入res
	if len(left) == 0 {
		res = append(res, right...)
	} else {
		res = append(res, left...)
	}
	return res
}
