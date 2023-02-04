package token

type TokenType string

type Token struct {
	Type    TokenType // 토큰의 타입 정보
	Literal string    // 토큰 자체의 문자값
}

const (
	ILLEGAL = "ILLEGAL" // 렉서가 어떤 토큰이나 문자를 해석할 수 없음
	EOF     = "EOF"     // 파일의 끝을 의미하며, 파서에게 보낼 중단 신호

	// 식별자와 리터럴
	IDENT = "IDENT" // 예. add, foobar, x, y, ...
	INT   = "INT"   // 예. 12343

	// 연산자
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// 구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 예약어(keywords)
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
