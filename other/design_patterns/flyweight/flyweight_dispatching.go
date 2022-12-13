package flyweight

import "bytes"

// TaxiPosition 出租车位置信息 x,y为外在数据信息，taxi为内在数据信息（享元对象）
type TaxiPosition struct {
	x    int
	y    int
	taxi *Taxi
}

func NewTaxiPosition(taxi *Taxi, x, y int) *TaxiPosition {
	return &TaxiPosition{
		taxi: taxi,
		x:    x,
		y:    y,
	}
}

// LocateFor 定位信息
func (p *TaxiPosition) LocateFor(monitorMap string) string {
	return p.taxi.LocateFor(monitorMap, p.x, p.y)
}

// TaxiDispatcher 出租车调度系统
type TaxiDispatcher struct {
	name   string
	traces map[string][]*TaxiPosition // 存储出租车当天轨迹信息，key为车牌号
}

func NewTaxiDispatcher(name string) *TaxiDispatcher {
	return &TaxiDispatcher{
		name:   name,
		traces: make(map[string][]*TaxiPosition),
	}
}

// AddTrace 添加轨迹
func (t *TaxiDispatcher) AddTrace(licensePlate, color, brand, company string, x, y int) {
	taxi := GetTaxiFactory().getTaxi(licensePlate, color, brand, company)
	t.traces[licensePlate] = append(t.traces[licensePlate], NewTaxiPosition(taxi, x, y))
}

// ShowTraces 显示轨迹
func (t *TaxiDispatcher) ShowTraces(licensePlate string) string {
	bytesBuf := bytes.Buffer{}
	for _, trace := range t.traces[licensePlate] {
		bytesBuf.WriteString(trace.LocateFor(t.name))
		bytesBuf.WriteByte('\n')
	}
	return bytesBuf.String()
}
