package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/adamerikoff/ponGo/src/lexer"
	"github.com/adamerikoff/ponGo/src/token"
)

const PROMPT = "|>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.NewLexer(line)

		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
