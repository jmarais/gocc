package autosuggest

import (
	"testing"

	"github.com/goccmack/gocc/example/autosuggest/ast"
	"github.com/goccmack/gocc/example/autosuggest/lexer"
	"github.com/goccmack/gocc/example/autosuggest/parser"
)

func TestQuery(t *testing.T) {
	query := `find author with "Seuss" at timerange 2019-01-01 13:13:12 -- 2019-01-01 13:13:13`
	lex := lexer.NewLexer([]byte(query))
	p := parser.NewParser()
	st, err := p.Parse(lex)
	if err != nil {
		panic(err)
	}
	s := st.(*ast.Search)
	t.Logf("Parsed: %#v:\n", s.FindAuthor)
}
