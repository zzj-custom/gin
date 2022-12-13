// Package factory 设计模式-工厂模式

// 工厂方法模式是一种创建型设计模式，其在父类中提供一个创建对象的方法， 允许子类决定实例化对象的类型。

package factory

// example
// 摊煎饼的小贩需要先摊个煎饼，再卖出去，摊煎饼就可以类比为一个工厂方法，根据顾客的喜好摊出不同口味的煎饼。

// Pancake 煎饼
type Pancake interface {
	// ShowFlour 煎饼使用的面粉
	ShowFlour() string
	// Value 煎饼价格
	Value() float32
}

// PancakeCook 煎饼厨师
type PancakeCook interface {
	// MakePancake 摊煎饼
	MakePancake() Pancake
}

// PancakeVendor 煎饼小贩
type PancakeVendor struct {
	PancakeCook
}

// NewPancakeVendor ...
func NewPancakeVendor(cook PancakeCook) *PancakeVendor {
	return &PancakeVendor{
		PancakeCook: cook,
	}
}

// SellPancake 卖煎饼，先摊煎饼，再卖
func (vendor *PancakeVendor) SellPancake() (money float32) {
	return vendor.MakePancake().Value()
}
