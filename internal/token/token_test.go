package token

import "testing"

func TestNew(t *testing.T) {
	var tok = New(LET, "let", 1)

	if tok.Type != LET {
		t.Fatalf("token.Type is not LET. got=%q", tok.Type)
	}

	if tok.Literal != "let" {
		t.Fatalf("token.Literal is not 'let'. got=%q", tok.Literal)
	}
}

func TestLookupIdent(t *testing.T) {
	var ident = "let"
	var tok = LookupIdent(ident)

	if tok != LET {
		t.Fatalf("token.Type is not LET. got=%q", tok)
	}

	ident = "variableName2"
	tok = LookupIdent(ident)

	if tok != IDENT {
		t.Fatalf("token.Type is not IDENT. got=%q", tok)
	}
}
