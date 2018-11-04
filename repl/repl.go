package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kitasuke/monkey-go/evaluator"
	"github.com/kitasuke/monkey-go/lexer"
	"github.com/kitasuke/monkey-go/object"
	"github.com/kitasuke/monkey-go/parser"
)

const Prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(Prompt)
		scanned := scanner.Scan()
		env := object.NewEnvironment()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
