package util

type BaseType interface {
	int | string | int8 | int16 | int32 | int64
}

func Add[T BaseType](slice []T, element []T, position int) []T {
	slice = append(slice, element...)
	copy(slice[position+len(element):], slice[position:])
	copy(slice[position:], element)
	return slice
}

func Delete[T BaseType](slice []T, count, position int) []T {
	return append(slice[:position], slice[count+position:]...)
}

func Filter[T BaseType](slice []T, fn func(x T) bool) []T {
	r := slice[:0]

	for _, s := range slice {
		if fn(s) {
			r = append(r, s)
		}
	}
	return r
}
