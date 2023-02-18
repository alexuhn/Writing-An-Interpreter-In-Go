package ast

import (
	"bytes"
	"monkey/token"
)

type Node interface {
	TokenLiteral() string // 디버깅 및 테스트용
	String() string
}

// 명령문(값을 만들지 않음)
type Statement interface {
	Node
	statementNode()
}

// 표현식(값을 만듦)
type Expression interface {
	Node
	expressionNode()
}

// Program 노드 = 모든 AST의 루트 노드
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Token token.Token // token.LET 토큰
	Name  *Identifier // 변수 바인딩 식별자
	Value Expression  // 값을 생성하는 표현식
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
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

type ReturnStatement struct {
	Token       token.Token // token.RETURN 토큰
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // 표현식의 첫 번째 토큰
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type Identifier struct {
	Token token.Token // token.IDENT 토큰
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type IntegerLiteral struct {
	Token token.Token // token.INT 토큰
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		// 각 명령문의 String 메서드를 호출하여 반환값을 버퍼에 씀
		out.WriteString(s.String())
	}

	return out.String()
}
