package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // 입력에서 현재 위치 즉, 현재 문자를 가리킴
	readPosition int  //입력에서 현재 읽는 위치 즉, 현재 문자의 다음을 가리킴
	ch           byte // 현재 조사하고 있는 문자 (참고. 유니코드가 아닌 ACSII 문자만 지원하기 때문에 byte 타입 사용 가능)
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// 렉서가 다음 문자로 이동하게 하는 메서드
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 끝에 도달한 경우 ASCII 코드 문자 "NUL"에 해당하는 0을 부여
		// 0 => 아무 것도 읽지 않은 상태 또는 파일의 끝(EOF)
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
