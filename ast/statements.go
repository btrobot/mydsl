package ast

import (
	"bytes"
	"github.com/btrobot/mydsl/token"
)

// LetStatement 表示变量声明语句
type LetStatement struct {
	Token token.Token // LET 词法单元
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) Position() (int, int) { return ls.Token.Line, ls.Token.Column }

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	
	out.WriteString(";")
	
	return out.String()
}

// ReturnStatement 表示返回语句
type ReturnStatement struct {
	Token       token.Token // RETURN 词法单元
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) Position() (int, int) { return rs.Token.Line, rs.Token.Column }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString(rs.TokenLiteral() + " ")
	
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	
	out.WriteString(";")
	
	return out.String()
}

// ExpressionStatement 表示表达式语句
type ExpressionStatement struct {
	Token      token.Token // 表达式的第一个词法单元
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) Position() (int, int) { return es.Token.Line, es.Token.Column }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String() + ";"
	}
	return ";"
}

// BlockStatement 表示代码块
type BlockStatement struct {
	Token      token.Token // { 词法单元
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) Position() (int, int) { return bs.Token.Line, bs.Token.Column }

func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString("{ ")
	
	for _, s := range bs.Statements {
		out.WriteString(s.String())
		out.WriteString(" ")
	}
	
	out.WriteString("}")
	
	return out.String()
}
