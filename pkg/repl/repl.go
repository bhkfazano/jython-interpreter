package repl

import (
	"bufio"
	"fmt"
	"io"
	"kogab-interpreter/internal/lexer"
	"kogab-interpreter/internal/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, user string) {
	var scanner = bufio.NewScanner(in)

	for {
		fmt.Printf("[%s] %s", user, PROMPT)
		var scanned = scanner.Scan()

		if !scanned {
			return
		}

		var line = scanner.Text()
		var lex = lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("[Jython] %+v\n", tok)
		}
	}
}
