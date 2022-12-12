// Package design_patterns 设计模式-解释器模式

// 解释器模式用于描述如何使用面向对象语言构成一个简单的语言解释器。在某些情况下，为了更好地描述某一些特定类型的问题，
// 我们可以创建一种新的语言，这种语言拥有自己的表达式和结构，即文法规则，这些问题的实例将对应为该语言中的句子。
// 此时，可以使用解释器模式来设计这种新的语言。对解释器模式的学习能够加深我们对面向对象思想的理解，并且掌握编程语言中文法规则的解释过程。

package interpreter

import "strings"

// example
// 定义一个解析特征值的语句解释器，提供是否包含特征值的终结表达式，并提供或表达式与且表达式，
// 同时，生成南极洲特征判断表达式，及美国人特征判断表达式，最后测试程序根据对象特征值描述，通过表达式判断是否为真。

// Expression 表达式接口，包含一个解释方法
type Expression interface {
	Interpret(context string) bool
}

// terminalExpression 终结符表达式，判断表达式中是否包含匹配数据
type terminalExpression struct {
	matchData string
}

func NewTerminalExpression(matchData string) *terminalExpression {
	return &terminalExpression{matchData: matchData}
}

// Interpret 判断是否包含匹配字符
func (t *terminalExpression) Interpret(context string) bool {
	if strings.Contains(context, t.matchData) {
		return true
	}
	return false
}

// orExpression 或表达式
type orExpression struct {
	left, right Expression
}

func NewOrExpression(left, right Expression) *orExpression {
	return &orExpression{
		left:  left,
		right: right,
	}
}

func (o *orExpression) Interpret(context string) bool {
	return o.left.Interpret(context) || o.right.Interpret(context)
}

// andExpression 与表达式
type andExpression struct {
	left, right Expression
}

func NewAndExpression(left, right Expression) *andExpression {
	return &andExpression{
		left:  left,
		right: right,
	}
}

func (o *andExpression) Interpret(context string) bool {
	return o.left.Interpret(context) && o.right.Interpret(context)
}
