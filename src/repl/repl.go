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
		fmt.Print(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.NewLexer(line)

		for tok := lexer.NextToken(); tok.Type != token.END_OF_FILE; tok = lexer.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
