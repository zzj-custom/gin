package builder

// Quantity 分量
type Quantity int

const (
	Small  Quantity = 1
	Middle Quantity = 5
	Large  Quantity = 10
)

type PancakeBuilder interface {
	// PutPaste 放面糊
	PutPaste(quantity Quantity)
	// PutEgg 放鸡蛋
	PutEgg(num int)
	// PutWafer 放薄脆
	PutWafer()
	// PutFlavour 放调料 Coriander香菜，Shallot葱 Sauce酱
	PutFlavour(hasCoriander, hasShallot, hasSauce bool)
	// Build 摊煎饼
	Build() *Pancake
}

// Pancake  煎饼
type Pancake struct {
	pasteQuantity Quantity // 面糊分量
	eggNum        int      // 鸡蛋数量
	wafer         string   // 薄脆
	hasCoriander  bool     // 是否放香菜
	hasShallot    bool     // 是否放葱
	hasSauce      bool     // 是否放酱
}
