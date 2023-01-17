package sort_arithmetic

// HeapSort 堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。
// 堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。
// 堆排序可以说是一种利用堆的概念来排序的选择排序。分为两种方法：
// 1. 大顶堆：每个节点的值都大于或等于其子节点的值，在堆排序算法中用于升序排列；arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
// 2. 小顶堆：每个节点的值都小于或等于其子节点的值，在堆排序算法中用于降序排列； arr[i] <= arr[2i+1] && arr[i] <= arr[2i+2]
// 堆排序的基本思想是：将待排序序列构造成一个大顶堆，此时，整个序列的最大值就是堆顶的根节点。
// 将其与末尾元素进行交换，此时末尾就为最大值。然后将剩余n-1个元素重新构造成一个堆，这样会得到n个元素的次小值。如此反复执行，便能得到一个有序序列了
func HeapSort(arr []int) []int {
	arrLen := len(arr)
	//初始化大顶堆
	BuildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		arrLen--
		heap(arr, 0, arrLen)
	}
	return arr
}

func BuildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heap(arr, i, arrLen)
	}
}

func heap(arr []int, i, arrLen int) {
	left := 2*i + 1  // 左子
	right := 2*i + 2 // 右子
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		// 交换数据
		swap1(arr, i, largest)
		// 调整二叉树
		heap(arr, largest, arrLen)
	}
}

func swap1(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
