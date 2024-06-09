package gogo_repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/adamerikoff/gogo/gogo_lexer"
	"github.com/adamerikoff/gogo/gogo_token"
)

const PROMPT = "=>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := gogo_lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != gogo_token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "| Type: %s, Literal: %s |\n", tok.Type, tok.Literal)
		}
	}
}
