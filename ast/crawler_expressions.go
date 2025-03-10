package ast

import (
	"bytes"
	"strings"
	"github.com/btrobot/mydsl/token"
)

// OpenExpression 表示网页打开操作
type OpenExpression struct {
	Token token.Token // OPEN 词法单元
	URL   Expression  // URL 表达式
}

func (oe *OpenExpression) expressionNode() {}
func (oe *OpenExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *OpenExpression) Position() (int, int) { return oe.Token.Line, oe.Token.Column }

func (oe *OpenExpression) String() string {
	var out bytes.Buffer
	
	out.WriteString("open(")
	if oe.URL != nil {
		out.WriteString(oe.URL.String())
	}
	out.WriteString(")")
	
	return out.String()
}

// ExtractExpression 表示数据提取操作
type ExtractExpression struct {
	Token    token.Token // EXTRACT 词法单元
	Source   Expression  // 源数据表达式
	Selector Expression  // 选择器表达式
}

func (ee *ExtractExpression) expressionNode() {}
func (ee *ExtractExpression) TokenLiteral() string { return ee.Token.Literal }
func (ee *ExtractExpression) Position() (int, int) { return ee.Token.Line, ee.Token.Column }

func (ee *ExtractExpression) String() string {
	var out bytes.Buffer
	
	out.WriteString("extract(")
	if ee.Source != nil {
		out.WriteString(ee.Source.String())
	}
	out.WriteString(", ")
	if ee.Selector != nil {
		out.WriteString(ee.Selector.String())
	}
	out.WriteString(")")
	
	return out.String()
}

// CollectExpression 表示数据收集操作
type CollectExpression struct {
	Token     token.Token // COLLECT 词法单元
	Source    Expression  // 源数据表达式
	Selectors []Expression // 选择器表达式列表
}

func (ce *CollectExpression) expressionNode() {}
func (ce *CollectExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CollectExpression) Position() (int, int) { return ce.Token.Line, ce.Token.Column }

func (ce *CollectExpression) String() string {
	var out bytes.Buffer
	
	out.WriteString("collect(")
	if ce.Source != nil {
		out.WriteString(ce.Source.String())
	}
	
	if len(ce.Selectors) > 0 {
		selectors := []string{}
		for _, sel := range ce.Selectors {
			selectors = append(selectors, sel.String())
		}
		out.WriteString(", ")
		out.WriteString(strings.Join(selectors, ", "))
	}
	
	out.WriteString(")")
	
	return out.String()
}

// AtExpression 表示选择器表达式
type AtExpression struct {
	Token    token.Token // AT 词法单元
	Selector Expression  // 选择器表达式
}

func (ae *AtExpression) expressionNode() {}
func (ae *AtExpression) TokenLiteral() string { return ae.Token.Literal }
func (ae *AtExpression) Position() (int, int) { return ae.Token.Line, ae.Token.Column }

func (ae *AtExpression) String() string {
	var out bytes.Buffer
	
	out.WriteString("@")
	if ae.Selector != nil {
		out.WriteString(ae.Selector.String())
	}
	
	return out.String()
}

// PipeExpression 表示管道操作
type PipeExpression struct {
	Token token.Token // PIPE 词法单元
	Left  Expression  // 左侧表达式
	Right Expression  // 右侧表达式
}

func (pe *PipeExpression) expressionNode() {}
func (pe *PipeExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PipeExpression) Position() (int, int) { return pe.Token.Line, pe.Token.Column }

func (pe *PipeExpression) String() string {
	var out bytes.Buffer
	
	out.WriteString("(")
	if pe.Left != nil {
		out.WriteString(pe.Left.String())
	}
	out.WriteString(" | ")
	if pe.Right != nil {
		out.WriteString(pe.Right.String())
	}
	out.WriteString(")")
	
	return out.String()
}
