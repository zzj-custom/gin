// Package builder 设计模式-生成器模式

// 生成器是一种创建型设计模式，使你能够分步骤创建复杂对象。
// 与其他创建型模式不同，生成器不要求产品拥有通用接口。这使得用相同地创建过程生成不同的产品成为可能。

package builder

// 还是摊煎饼的例子，摊煎饼分为四个步骤，1放面糊、2放鸡蛋、3放调料、4放薄脆， 通过四个创建过程，制作好一个煎饼，
// 这个摊煎饼的过程就好比煎饼生成器接口，不同生成器的实现就相当于摊不同品类的煎饼，
// 比如正常的煎饼，健康的煎饼（可能用的是粗粮面、柴鸡蛋、非油炸薄脆、不放酱等），
// 生成器接口方法也可以通过参数控制煎饼的大小， 比如放两勺面糊，放2个鸡蛋等。
// 生成器的使用者为了避免每次都调用相同的构建步骤，也可以通过包装类固定几种构建过程，生成几类常用的产品，
// 就好像摊煎饼有几类常卖固定成品，比如普通的，加两个鸡蛋的，不要香菜的等等，这几类固定构建过程提前定制好，
// 直接通过简单工厂方法就直接创建，如果用户再需要细粒度的定制构建，再通过生成器创建。

type normalPancakeBuilder struct {
	pasteQuantity Quantity // 面糊量
	eggNum        int      // 鸡蛋数量
	friedWafer    string   // 油炸薄脆
	hasCoriander  bool     // 是否放香菜
	hasShallot    bool     // 是否放葱
	hasHotSauce   bool     // 是否放辣味酱
}

func NewNormalPancakeBuilder() *normalPancakeBuilder {
	return &normalPancakeBuilder{}
}

func (n *normalPancakeBuilder) PutPaste(quantity Quantity) {
	n.pasteQuantity = quantity
}

func (n *normalPancakeBuilder) PutEgg(num int) {
	n.eggNum = num
}

func (n *normalPancakeBuilder) PutWafer() {
	n.friedWafer = "油炸的薄脆"
}

func (n *normalPancakeBuilder) PutFlavour(hasCoriander, hasShallot, hasSauce bool) {
	n.hasCoriander = hasCoriander
	n.hasShallot = hasShallot
	n.hasHotSauce = hasSauce
}

func (n *normalPancakeBuilder) Build() *Pancake {
	return &Pancake{
		pasteQuantity: n.pasteQuantity,
		eggNum:        n.eggNum,
		wafer:         n.friedWafer,
		hasCoriander:  n.hasCoriander,
		hasShallot:    n.hasShallot,
		hasSauce:      n.hasHotSauce,
	}
}

type healthyPancakeBuilder struct {
	milletPasteQuantity Quantity // 小米面糊量
	chaiEggNum          int      // 柴鸡蛋数量
	nonFriedWafer       string   // 非油炸薄脆
	hasCoriander        bool     // 是否放香菜
	hasShallot          bool     // 是否放葱
}

func NewHealthyPancakeBuilder() *healthyPancakeBuilder {
	return &healthyPancakeBuilder{}
}

func (n *healthyPancakeBuilder) PutPaste(quantity Quantity) {
	n.milletPasteQuantity = quantity
}

func (n *healthyPancakeBuilder) PutEgg(num int) {
	n.chaiEggNum = num
}

func (n *healthyPancakeBuilder) PutWafer() {
	n.nonFriedWafer = "非油炸的薄脆"
}

func (n *healthyPancakeBuilder) PutFlavour(hasCoriander, hasShallot, _ bool) {
	n.hasCoriander = hasCoriander
	n.hasShallot = hasShallot
}

func (n *healthyPancakeBuilder) Build() *Pancake {
	return &Pancake{
		pasteQuantity: n.milletPasteQuantity,
		eggNum:        n.chaiEggNum,
		wafer:         n.nonFriedWafer,
		hasCoriander:  n.hasCoriander,
		hasShallot:    n.hasShallot,
		hasSauce:      false,
	}
}

// PancakeCook 摊煎饼师傅
type PancakeCook struct {
	builder PancakeBuilder
}

func NewPancakeCook(builder PancakeBuilder) *PancakeCook {
	return &PancakeCook{
		builder: builder,
	}
}

// SetPancakeBuilder 重新设置煎饼构造器
func (p *PancakeCook) SetPancakeBuilder(builder PancakeBuilder) {
	p.builder = builder
}

// MakePancake 摊一个一般煎饼
func (p *PancakeCook) MakePancake() *Pancake {
	p.builder.PutPaste(Middle)
	p.builder.PutEgg(1)
	p.builder.PutWafer()
	p.builder.PutFlavour(true, true, true)
	return p.builder.Build()
}

// MakeBigPancake 摊一个巨无霸煎饼
func (p *PancakeCook) MakeBigPancake() *Pancake {
	p.builder.PutPaste(Large)
	p.builder.PutEgg(3)
	p.builder.PutWafer()
	p.builder.PutFlavour(true, true, true)
	return p.builder.Build()
}

// MakePancakeForFlavour 摊一个自选调料霸煎饼
func (p *PancakeCook) MakePancakeForFlavour(hasCoriander, hasShallot, hasSauce bool) *Pancake {
	p.builder.PutPaste(Large)
	p.builder.PutEgg(3)
	p.builder.PutWafer()
	p.builder.PutFlavour(hasCoriander, hasShallot, hasSauce)
	return p.builder.Build()
}
