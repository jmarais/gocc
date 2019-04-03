package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	parseError "github.com/goccmack/gocc/example/autosuggest/errors"
	"github.com/goccmack/gocc/example/autosuggest/lexer"
	"github.com/goccmack/gocc/example/autosuggest/parser"
)

// small demo of how the suggestion parser can help a user
// without knowledge of the ebnf definition.
// Start typing your query and press `enter` to submit and get a list
// of suggestions back.
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter query (press enter to submit):\n")
	query := ""
	p := parser.NewParser()
	for {
		s, _ := reader.ReadString('\n')
		query += s
		query = strings.TrimRight(query, "\n")
		l := lexer.NewLexer([]byte(query))

		attribs, err := p.SuggestParse(l)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		if len(attribs) > 0 {
			switch t := attribs[len(attribs)-1].(type) {
			case *parseError.Error:
				fmt.Printf("ExpectedTokens:\n")
				for i, s := range t.ExpectedTokens {
					fmt.Printf("[%d] %v\n", i, s)
				}
			default:
				fmt.Printf("No suggestions\n")
			}
		}
		fmt.Printf(query)
	}
}
