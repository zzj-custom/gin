// Package design_patterns 设计模式-访问者模式

// 访问者模式是一种行为设计模式，它能将算法与其所作用的对象隔离开来。允许你在不修改已有代码的情况下向已有类层次结构中增加新的行为。
// 访问者接口需要根据被访问者具体类，定义多个相似的访问方法，每个具体类对应一个访问方法；
// 每个被访问者需要实现一个接受访问者对象的方法，方法的实现就是去调用访问者接口对应该类的访问方法；
// 这个接受方法可以传入不同目的访问者接口的具体实现，从而在不修改被访问对象的前提下，增加新的功能；

package visitor

import "fmt"

// example
// 公司中存在多种类型的员工，包括产品经理、软件工程师、人力资源等，他们的KPI指标不尽相同，
// 产品经理为上线产品数量及满意度，
// 软件工程师为实现的需求数及修改bug数，
// 人力资源为招聘员工的数量；
// 公司要根据员工完成的KPI进行表彰公示，同时根据KPI完成情况定薪酬，
// 这些功能都是员工类职责之外的，不能修改员工本身的类，我们通过访问者模式，实现KPI表彰排名及薪酬发放；

// Employee 员工接口
type Employee interface {
	KPI() string                    // 完成kpi信息
	Accept(visitor EmployeeVisitor) // 接受访问者对象
}

// productManager 产品经理
type productManager struct {
	name         string // 名称
	productNum   int    // 上线产品数
	satisfaction int    // 平均满意度
}

func NewProductManager(name string, productNum int, satisfaction int) *productManager {
	return &productManager{
		name:         name,
		productNum:   productNum,
		satisfaction: satisfaction,
	}
}

func (p *productManager) KPI() string {
	return fmt.Sprintf("产品经理%s，上线%d个产品，平均满意度为%d", p.name, p.productNum, p.satisfaction)
}

func (p *productManager) Accept(visitor EmployeeVisitor) {
	visitor.VisitProductManager(p)
}

// softwareEngineer 软件工程师
type softwareEngineer struct {
	name           string // 姓名
	requirementNum int    // 完成需求数
	bugNum         int    // 修复问题数
}

func NewSoftwareEngineer(name string, requirementNum int, bugNum int) *softwareEngineer {
	return &softwareEngineer{
		name:           name,
		requirementNum: requirementNum,
		bugNum:         bugNum,
	}
}

func (s *softwareEngineer) KPI() string {
	return fmt.Sprintf("软件工程师%s，完成%d个需求，修复%d个问题", s.name, s.requirementNum, s.bugNum)
}

func (s *softwareEngineer) Accept(visitor EmployeeVisitor) {
	visitor.VisitSoftwareEngineer(s)
}

// hr 人力资源
type hr struct {
	name       string // 姓名
	recruitNum int    // 招聘人数
}

func NewHR(name string, recruitNum int) *hr {
	return &hr{
		name:       name,
		recruitNum: recruitNum,
	}
}

func (h *hr) KPI() string {
	return fmt.Sprintf("人力资源%s，招聘%d名员工", h.name, h.recruitNum)
}

func (h *hr) Accept(visitor EmployeeVisitor) {
	visitor.VisitHR(h)
}
