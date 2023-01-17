package sort_arithmetic

func swap(list []int, i, j int) []int {
	list[i], list[j] = list[j], list[i]
	return list
}
