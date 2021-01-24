package repl

import (
	"bufio"
	"github.com/keigodasu/writing-an-interpreter-in-go/lexer"
	"github.com/keigodasu/writing-an-interpreter-in-go/token"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
