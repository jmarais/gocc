package parser

import (
	parseError "github.com/goccmack/gocc/example/autosuggest/errors"
	"github.com/goccmack/gocc/example/autosuggest/token"
)

type suggestScanner struct {
	Scanner
}

func (s *suggestScanner) Scan() (tok *token.Token) {
	nextToken := s.Scanner.Scan()
	if nextToken.Type == token.EOF {
		return &token.Token{Type: token.INVALID}
	}
	return nextToken

}

func (p *Parser) SuggestParse(scanner Scanner) (atribs []Attrib, err error) {
	_, err = p.Parse(&suggestScanner{scanner})
	if err != nil {
		if e, ok := err.(*parseError.Error); ok {
			if e.Err == nil {
				p.stack.push(0, e)
			}
		} else {
			return nil, err
		}
	}
	return p.stack.attrib, nil
}
