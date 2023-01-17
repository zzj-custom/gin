package sort_arithmetic

// ShellSort 希尔排序
// message: 先将整个待排序列分割成若干个子序列，对若个子序列分别进行插入排序，待整个待排序列基本有序时，对整体进行插入排序。
// 1. 选择一个增量序列t1，t2，…，tk，其中ti>ti+1，tk=1；
// 2. 按增量序列个数k，对序列进行k 趟排序；
// 3. 每趟排序，根据对应的增量ti，将待排序列分割成若干个长度为 m 的子序列，分别对各子表进行直接插入排序。
// 仅增量因子为1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
func ShellSort(arr []int) []int {
	length := len(arr)
	for gap := length / 2; gap >= 1; gap = gap / 2 {
		// 缩小增量序列，希尔建议每次缩小一半
		for i := gap; i < length; i++ {
			// 子序列
			temp := arr[i]
			j := i - gap
			for ; j >= 0 && temp < arr[j]; j = j - gap {
				arr[j+gap] = arr[j]
			}
			arr[j+gap] = temp
		}
	}
	return arr
}
