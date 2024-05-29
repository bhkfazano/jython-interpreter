package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	BANG     = "BANG"
	ASTERISK = "ASTERISK"
	SLASH    = "SLASH"

	LT  = "LT"
	GT  = "GT"
	LTE = "LTE"
	GTE = "GTE"

	EQ     = "EQ"
	NOT_EQ = "NOT_EQ"

	// Delimiters
	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"

	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
	RBRACE = "RBRACE"

	// Keywords
	FUN    = "FUN"
	CONST  = "CONST"
	LET    = "LET"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

var keywords = map[string]TokenType{
	"fun":    FUN,
	"const":  CONST,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func New(tokenType TokenType, literal string, line int) Token {
	return Token{Type: tokenType, Literal: literal, Line: line}
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
