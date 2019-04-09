package main

import (
	// "bufio"
	"fmt"
	parseError "github.com/goccmack/gocc/example/autosuggest/errors"
	"github.com/goccmack/gocc/example/autosuggest/lexer"
	"github.com/goccmack/gocc/example/autosuggest/parser"
	"os"
	"os/exec"
)

// small demo of how the suggestion parser can help a user
// without knowledge of the ebnf definition.
// Start typing your query and press `enter` to submit and get a list
// of suggestions back.
func main() {
	// reader := bufio.NewReader(os.Stdin)
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	buff := make([]byte, 1)
	fmt.Printf("Enter query (press enter to submit):\n")
	query := make([]byte, 0)
	p := parser.NewParser()
	for {
		os.Stdin.Read(buff)
		// n, _ := reader.Read(buff)
		if buff[0] == byte(0x7f) {
			if len(query) > 1 {
				query = query[:len(query)-1]
			} else {
				query = query[0:0]
			}
		} else {
			query = append(query, buff...)
		}

		l := lexer.NewLexer(query)

		attribs, err := p.SuggestParse(l)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		if len(attribs) > 0 {
			switch t := attribs[len(attribs)-1].(type) {
			case *parseError.Error:
				fmt.Printf("\nExpectedTokens:\n")
				for i, s := range t.ExpectedTokens {
					fmt.Printf("[%d] %v\n", i, s)
				}
			default:
				fmt.Printf("No suggestions\n")
			}
		}
		fmt.Printf("%s", string(query))
	}
}
