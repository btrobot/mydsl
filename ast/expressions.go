package ast

import (
	"bytes"
	"fmt"
	"strings"
	"github.com/btrobot/mydsl/token"
)

// Identifier 表示标识符
type Identifier struct {
	Token token.Token // IDENT 词法单元
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) Position() (int, int) { return i.Token.Line, i.Token.Column }
func (i *Identifier) String() string { return i.Value }

// IntegerLiteral 表示整数字面量
type IntegerLiteral struct {
	Token token.Token // INT 词法单元
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) Position() (int, int) { return il.Token.Line, il.Token.Column }
func (il *IntegerLiteral) String() string { return il.Token.Literal }

// FloatLiteral 表示浮点数字面量
type FloatLiteral struct {
	Token token.Token // FLOAT 词法单元
	Value float64
}

func (fl *FloatLiteral) expressionNode() {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) Position() (int, int) { return fl.Token.Line, fl.Token.Column }
func (fl *FloatLiteral) String() string { return fl.Token.Literal }

// StringLiteral 表示字符串字面量
type StringLiteral struct {
	Token token.Token // STRING 词法单元
	Value string
}

func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) Position() (int, int) { return sl.Token.Line, sl.Token.Column }
func (sl *StringLiteral) String() string { return fmt.Sprintf("%q", sl.Value) }

// BooleanLiteral 表示布尔字面量
type BooleanLiteral struct {
	Token token.Token // TRUE 或 FALSE 词法单元
	Value bool
}

func (bl *BooleanLiteral) expressionNode() {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BooleanLiteral) Position() (int, int) { return bl.Token.Line, bl.Token.Column }
func (bl *BooleanLiteral) String() string { return bl.Token.Literal }

// NullLiteral 表示 null 字面量
type NullLiteral struct {
	Token token.Token // NULL 词法单元
}

func (nl *NullLiteral) expressionNode() {}
func (nl *NullLiteral) TokenLiteral() string { return nl.Token.Literal }
func (nl *NullLiteral) Position() (int, int) { return nl.Token.Line, nl.Token.Column }
func (nl *NullLiteral) String() string { return "null" }

// PrefixExpression 表示前缀表达式
type PrefixExpression struct {
	Token    token.Token // 前缀运算符词法单元
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) Position() (int, int) { return pe.Token.Line, pe.Token.Column }

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	
	return out.String()
}

// InfixExpression 表示中缀表达式
type InfixExpression struct {
	Token    token.Token // 运算符词法单元
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) Position() (int, int) { return ie.Token.Line, ie.Token.Column }

func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	
	return out.String()
}

// ArrayLiteral 表示数组字面量
type ArrayLiteral struct {
	Token    token.Token // [ 词法单元
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) Position() (int, int) { return al.Token.Line, al.Token.Column }

func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	
	return out.String()
}

// IndexExpression 表示索引表达式
type IndexExpression struct {
	Token token.Token // [ 词法单元
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) Position() (int, int) { return ie.Token.Line, ie.Token.Column }

func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	
	return out.String()
}

// ObjectLiteral 表示对象字面量
type ObjectLiteral struct {
	Token token.Token // { 词法单元
	Pairs map[Expression]Expression
}

func (ol *ObjectLiteral) expressionNode() {}
func (ol *ObjectLiteral) TokenLiteral() string { return ol.Token.Literal }
func (ol *ObjectLiteral) Position() (int, int) { return ol.Token.Line, ol.Token.Column }

func (ol *ObjectLiteral) String() string {
	var out bytes.Buffer
	
	pairs := []string{}
	for key, value := range ol.Pairs {
		pairs = append(pairs, key.String()+": "+value.String())
	}
	
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	
	return out.String()
}
