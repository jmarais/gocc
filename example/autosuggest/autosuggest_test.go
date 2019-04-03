package autosuggest

import (
	"testing"

	"github.com/goccmack/gocc/example/autosuggest/lexer"
	"github.com/goccmack/gocc/example/autosuggest/parser"
)

// The default parser will cause an error with incomplete input.
// This example is missing the date-range after the 'timerange' token,
// however the default parser will only inform you of the next token it
// is expecting:
// "expected one of: int_lit dayMonth number space"
// As an end user of this grammar, I might not be the one who wrote the bnf
// so the error message might not make sense to me.
// I might be unsure what it means with 'dayMonth' or even 'int_lit' if I am
// not familiar with the defined ebnf productions.
func TestDefaultParserError(t *testing.T) {
	query := `textsearch "The cat in the hat" with "It was too wet to play" at timerange`
	lex := lexer.NewLexer([]byte(query))
	p := parser.NewParser()
	_, err := p.Parse(lex)
	t.Logf("Parse Error: %s:\n", err)
}

// Using the suggestionParser you can get all the tokens on the stack as well
// as the next suggested token list.
// All the tokens on the stack is returned if the user wants to additional
// processing on the list of parsed tokens. The suggestions is available at
// the end of the list in the form of an errorToken (this should be improved,
// perhaps a suggestion token type>).
func TestSuggestionParser(t *testing.T) {
	// query := `find author with "Seuss" at timerange 2019-01-01 13:13:12 -- 2019-01-01 13:13:13`
	query := `textsearch "The cat in the hat" wit`
	lex := lexer.NewLexer([]byte(query))
	p := parser.NewParser()
	attribs, err := p.SuggestParse(lex)
	if err != nil {
		t.Logf("Parse Error: %s:\n", err)
		return
	}
	for i, attrib := range attribs {
		t.Logf("[%d]: %#v:\n", i, attrib)
	}
}
