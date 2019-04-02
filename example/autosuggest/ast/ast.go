package ast

import (
	"github.com/goccmack/gocc/example/autosuggest/token"
	"time"
)

type Search struct {
	FindTitle  *FindTitle
	FindAuthor *FindAuthor
	TextSearch *TextSearch
}

type FindAuthor struct {
	Find
}
type FindTitle struct {
	Find
}
type Find struct {
	Criteria *Criteria
	Interval *DateTimeInterval
}
type TextSearch struct{}
type Criteria struct{}
type DateTimeInterval struct{}
type Book struct{}
type Interval struct{}
type DateTime struct{}
type Space struct {
	s string
}

func NewFindAuthor(criteria, dateInterval interface{}) *FindAuthor {
	// findType *FindType, crit *Criteria, interval *DateTimeInterval
	// paramStr := ""
	// var err error
	// if param != nil {
	// 	plit := param.(*token.Token).Lit
	// 	paramStr, err = strconv.Unquote(string(plit))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// return &UDF{Function: newString(name), Parameter: paramStr}, nil
	return &FindAuthor{Find{}}
}

func NewFindTitle(criteria, dateInterval interface{}) *FindTitle {
	return &FindTitle{Find{}}
}
func NewTextSearch(book, criteria, dateInterval interface{}) *TextSearch {
	// book *Book, crit *Criteria, interval *DateTimeInterval
	return &TextSearch{}
}
func NewInterval(from, to interface{}) *DateTimeInterval {
	// from, to *DateTime
	return &DateTimeInterval{}
}
func NewTime(date, format string) *time.Time {
	// date, format string
	return &time.Time{}
}

func NewString(s interface{}) string {
	return string(s.(*token.Token).Lit)
}

func NewSpace(s interface{}) *Space {
	return &Space{string(s.(*token.Token).Lit)}
}
