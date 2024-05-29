package lexer

import (
	"kogab-interpreter/internal/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	var input string = `
let five = 5;
let ten = 10;
let add = fun(x, y) {
	x + y;
};

let result = add(five, ten);

if (result != 15) {
	return false;
} else {
	result = result / 2;
	result = result * 3;
	return true;
}

10 == 10;
10 != 9;
10 < 9;
10 > 9;
10 <= 10;
10 >= 10;

let reversedBool = !true;
let anotherBool = !false;
`

	var tests = []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUN, "fun"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "result"},
		{token.NOT_EQ, "!="},
		{token.INT, "15"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "result"},
		{token.SLASH, "/"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "result"},
		{token.ASTERISK, "*"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.LT, "<"},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.LTE, "<="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.GTE, ">="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "reversedBool"},
		{token.ASSIGN, "="},
		{token.BANG, "!"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "anotherBool"},
		{token.ASSIGN, "="},
		{token.BANG, "!"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	var l = New(input)

	for i, tt := range tests {
		var tok = l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
