package ast

import (
	"time"

	"github.com/goccmack/gocc/example/autosuggest/token"
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
type Criteria struct {
	s string
}
type DateTimeInterval struct{}
type Book struct{}
type Interval struct{}
type DateTime struct{}
type Space struct {
	s string
}

func NewFindAuthor(criteria string, dateInterval interface{}) *FindAuthor {
	var dateInter *DateTimeInterval
	if dateInterval != nil {
		dateInter = dateInterval.(*DateTimeInterval)
	}
	return &FindAuthor{Find{Criteria: &Criteria{criteria}, Interval: dateInter}}
}

func NewFindTitle(criteria string, dateInterval interface{}) *FindTitle {
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
func NewTime(date, format interface{}) *time.Time {
	// date, format string
	return &time.Time{}
}

func NewString(s interface{}) string {
	return string(s.(*token.Token).Lit)
}

func NewSpace(s interface{}) *Space {
	return &Space{string(s.(*token.Token).Lit)}
}
