package factory

// cornPancakeCook 制作玉米面煎饼厨师
type cornPancakeCook struct{}

func NewCornPancakeCook() *cornPancakeCook {
	return &cornPancakeCook{}
}

func (cook *cornPancakeCook) MakePancake() Pancake {
	return NewCornPancake()
}

// milletPancakeCook 制作小米面煎饼厨师
type milletPancakeCook struct{}

func NewMilletPancakeCook() *milletPancakeCook {
	return &milletPancakeCook{}
}

func (cook *milletPancakeCook) MakePancake() Pancake {
	return NewMilletPancake()
}

// cornPancake 玉米面煎饼
type cornPancake struct{}

// NewCornPancake ...
func NewCornPancake() *cornPancake {
	return &cornPancake{}
}

func (cake *cornPancake) ShowFlour() string {
	return "玉米面"
}

func (cake *cornPancake) Value() float32 {
	return 5.0
}

// milletPancake 小米面煎饼
type milletPancake struct{}

func NewMilletPancake() *milletPancake {
	return &milletPancake{}
}

func (cake *milletPancake) ShowFlour() string {
	return "小米面"
}

func (cake *milletPancake) Value() float32 {
	return 8.0
}
