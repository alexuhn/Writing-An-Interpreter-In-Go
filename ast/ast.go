package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string // 디버깅 및 테스트용
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

type Identifier struct {
	Token token.Token // token.IDENT 토큰
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
