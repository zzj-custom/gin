package composite

type Region interface {
	Name() string
	Population() int
	GDP() float64
}
