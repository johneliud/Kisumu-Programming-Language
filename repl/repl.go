package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/johneliud/Kisumu-Programming-Language/lexer"
	"github.com/johneliud/Kisumu-Programming-Language/token"
)

const PROMPT = "->"

func Start(input io.Reader, output io.Writer) {
	/*
	   ksmFile, err := os.Open("./source-file.ksm")
	   if err != nil {
	   fmt.Printf("Failed to open source file: %s\n", err)
	   os.Exit(1)
	   }
	   defer ksmFile.Close()
	*/
	scanner := bufio.NewScanner(input)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		// Loops through each token produced by the lexer until EOF is reached
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}
