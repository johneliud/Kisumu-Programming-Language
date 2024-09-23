package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/johneliud/Kisumu-Programming-Language/evaluator"
	"github.com/johneliud/Kisumu-Programming-Language/lexer"
	"github.com/johneliud/Kisumu-Programming-Language/object"
	"github.com/johneliud/Kisumu-Programming-Language/parser"
)

func Start(output io.Writer) {
	sourceFile, err := os.Open("source-code.ksm")
	if err != nil {
		fmt.Println("Error opening file!", err)
		return
	}
	defer sourceFile.Close()

	scanner := bufio.NewScanner(sourceFile)
	env := object.NewEnvironment()

	for {
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(output, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(output, evaluated.Inspect())
			io.WriteString(output, "\n")
		}
	}
}

func printParserErrors(output io.Writer, errors []string) {
	io.WriteString(output, "An unexpected error just occurred\n")
	io.WriteString(output, " parser errors:\n")

	for _, msg := range errors {
		io.WriteString(output, "\t"+msg+"\n")
	}
}
