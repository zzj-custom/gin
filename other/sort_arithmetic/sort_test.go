package sort_arithmetic

import (
	"github.com/go-playground/assert/v2"
	"github.com/gookit/goutil/dump"
	"testing"
)

func Test_UtilsSwap(t *testing.T) {
	list := []int{1, 2, 3, 4}
	want := []int{1, 4, 3, 2}
	assert.Equal(t, swap(list, 1, 3), want)
}

func TestBubbleSort1(t *testing.T) {
	list := []int{1, 4, 3, 2}
	dump.P(BubbleSort1(list))
}

func TestBubbleSort(t *testing.T) {
	list := []int{5, 1, 4, 3, 2}
	want := []int{1, 2, 3, 4, 5}
	assert.Equal(t, SelectionSort(list), want)
}

func TestInsertionSort(t *testing.T) {
	list := []int{5, 1, 4, 3, 2}
	want := []int{1, 2, 3, 4, 5}
	assert.Equal(t, InsertionSort(list), want)
}

func TestShellSort(t *testing.T) {
	list := []int{5, 1, 4, 3, 2}
	want := []int{1, 2, 3, 4, 5}
	assert.Equal(t, ShellSort(list), want)
}

func TestMergeSort(t *testing.T) {
	list := []int{5, 1, 4, 3, 2}
	want := []int{1, 2, 3, 4, 5}
	assert.Equal(t, MergeSort(list), want)
}

func TestQuickSort(t *testing.T) {
	list := []int{5, 1, 4, 3, 2, 6, 9, 12, 7, 44, 22, 33, 56, 18}
	want := []int{1, 2, 3, 4, 5, 6, 7, 9, 12, 18, 22, 33, 44, 56}
	assert.Equal(t, QuickSort(list), want)
}

func TestHeapSort(t *testing.T) {
	list := []int{5, 1, 4, 3, 2, 6, 9, 12, 7, 44, 22, 33, 56, 18}
	want := []int{1, 2, 3, 4, 5, 6, 7, 9, 12, 18, 22, 33, 44, 56}
	assert.Equal(t, HeapSort(list), want)
}

func TestCountingSort(t *testing.T) {
	list := []int{5, 1, 4, 3, 2, 6, 9, 12, 7, 44, 22, 33, 56, 18}
	want := []int{1, 2, 3, 4, 5, 6, 7, 9, 12, 18, 22, 33, 44, 56}
	assert.Equal(t, CountingSort(list), want)
}

func TestBucketSort(t *testing.T) {
	list := []int{5, 1, 4, 3, 2, 6, 9, 12, 7, 44, 22, 33, 56, 18}
	want := []int{1, 2, 3, 4, 5, 6, 7, 9, 12, 18, 22, 33, 44, 56}
	assert.Equal(t, BucketSort(list, 5), want)
}
