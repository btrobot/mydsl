package ast

import (
	"bytes"
	"github.com/btrobot/mydsl/token"
)

// Node 表示 AST 中的节点
type Node interface {
	TokenLiteral() string // 返回与节点关联的词法单元的字面值
	String() string       // 返回节点的字符串表示
	Position() (int, int) // 返回节点的位置（行号和列号）
}

// Statement 表示语句
type Statement interface {
	Node
	statementNode() // 标记方法，用于区分语句和表达式
}

// Expression 表示表达式
type Expression interface {
	Node
	expressionNode() // 标记方法，用于区分语句和表达式
}

// Program 表示整个程序，是 AST 的根节点
type Program struct {
	Statements []Statement
}

// TokenLiteral 返回程序的第一个语句的词法单元字面值
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String 返回程序的字符串表示
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Position 返回程序的位置（第一个语句的位置）
func (p *Program) Position() (int, int) {
	if len(p.Statements) > 0 {
		return p.Statements[0].Position()
	}
	return 0, 0
}
