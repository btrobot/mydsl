package ast

import (
	"bytes"
	"github.com/btrobot/mydsl/token"
)

// IfExpression 表示条件表达式
type IfExpression struct {
	Token       token.Token // IF 词法单元
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) Position() (int, int) { return ie.Token.Line, ie.Token.Column }

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// WhileExpression 表示循环表达式
type WhileExpression struct {
	Token     token.Token // WHILE 词法单元
	Condition Expression
	Body      *BlockStatement
}

func (we *WhileExpression) expressionNode() {}
func (we *WhileExpression) TokenLiteral() string { return we.Token.Literal }
func (we *WhileExpression) Position() (int, int) { return we.Token.Line, we.Token.Column }

func (we *WhileExpression) String() string {
	var out bytes.Buffer

	out.WriteString("while ")
	out.WriteString(we.Condition.String())
	out.WriteString(" ")
	out.WriteString(we.Body.String())

	return out.String()
}

// ForExpression 表示 for 循环表达式
type ForExpression struct {
	Token      token.Token // FOR 词法单元
	Identifier *Identifier
	Iterable   Expression
	Body       *BlockStatement
}

func (fe *ForExpression) expressionNode() {}
func (fe *ForExpression) TokenLiteral() string { return fe.Token.Literal }
func (fe *ForExpression) Position() (int, int) { return fe.Token.Line, fe.Token.Column }

func (fe *ForExpression) String() string {
	var out bytes.Buffer

	out.WriteString("for ")
	out.WriteString(fe.Identifier.String())
	out.WriteString(" in ")
	out.WriteString(fe.Iterable.String())
	out.WriteString(" ")
	out.WriteString(fe.Body.String())

	return out.String()
}

// FunctionLiteral 表示函数字面量
type FunctionLiteral struct {
	Token      token.Token // FUNCTION 词法单元
	Parameters []*Identifier
	Body       *BlockStatement
	Name       string // 可选的函数名
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) Position() (int, int) { return fl.Token.Line, fl.Token.Column }

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	if fl.Name != "" {
		out.WriteString(" ")
		out.WriteString(fl.Name)
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// CallExpression 表示函数调用表达式
type CallExpression struct {
	Token     token.Token // ( 词法单元
	Function  Expression  // 标识符或函数字面量
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) Position() (int, int) { return ce.Token.Line, ce.Token.Column }

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
