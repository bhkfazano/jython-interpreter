package lexer

import (
	"fmt"
	"kogab-interpreter/internal/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int
}

func New(input string) *Lexer {
	var lexer = Lexer{input: input, line: 1}
	lexer.readChar()
	return &lexer
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	if string(l.ch) == "\n" {
		l.line++
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = token.New(token.PLUS, string(l.ch), l.line)
	case '-':
		tok = token.New(token.MINUS, string(l.ch), l.line)
	case '*':
		tok = token.New(token.ASTERISK, string(l.ch), l.line)
	case ';':
		tok = token.New(token.SEMICOLON, string(l.ch), l.line)
	case '(':
		tok = token.New(token.LPAREN, string(l.ch), l.line)
	case ')':
		tok = token.New(token.RPAREN, string(l.ch), l.line)
	case ',':
		tok = token.New(token.COMMA, string(l.ch), l.line)
	case '{':
		tok = token.New(token.LBRACE, string(l.ch), l.line)
	case '}':
		tok = token.New(token.RBRACE, string(l.ch), l.line)
	case '/':
		if l.peekChar() == '/' {
			l.skipComment()
			return l.NextToken()
		}
		tok = token.New(token.SLASH, string(l.ch), l.line)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
		tok.Line = l.line
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.EQ, string(ch)+string(l.ch), l.line)
		} else {
			tok = token.New(token.ASSIGN, string(l.ch), l.line)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.NOT_EQ, string(ch)+string(l.ch), l.line)
		} else {
			tok = token.New(token.BANG, string(l.ch), l.line)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.LTE, string(ch)+string(l.ch), l.line)
		} else {
			tok = token.New(token.LT, string(l.ch), l.line)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.GTE, string(ch)+string(l.ch), l.line)
		} else {
			tok = token.New(token.GT, string(l.ch), l.line)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isDigit(l.ch) {
			return l.readNumber()
		} else if isAlpha(l.ch) {
			return l.readIdentifier()
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch), l.line)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readNumber() token.Token {
	var position = l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	var scannedNum = l.input[position:l.position]
	return token.New(token.INT, scannedNum, l.line)
}

func (l *Lexer) readIdentifier() token.Token {
	var position = l.position
	for isAlphanumeric(l.ch) {
		l.readChar()
	}
	var ident = l.input[position:l.position]
	return token.New(token.LookupIdent(ident), ident, l.line)
}

func (l *Lexer) readString() string {
	l.readChar()
	var position = l.position
	for l.ch != '"' {
		if l.ch == 0 {
			panic("[ERROR] Unexpected EOF while reading string on line " + fmt.Sprint(l.line))
		}
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) HasNext() bool {
	return l.position < len(l.input)
}

func isAlpha(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAlphanumeric(ch byte) bool {
	return isAlpha(ch) || isDigit(ch)
}

func (l *Lexer) skipWhitespace() {
	for string(l.ch) == " " || string(l.ch) == "\t" || string(l.ch) == "\n" || string(l.ch) == "\r" {
		l.readChar()
	}
}
