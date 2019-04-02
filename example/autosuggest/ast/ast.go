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

type TextSearch struct {
	Book     string
	Criteria *Criteria
	Interval *DateTimeInterval
}

type Criteria struct {
	S string
}

type DateTimeInterval struct {
	From time.Time
	To   time.Time
}

type Space struct {
	S string
}

func NewFindAuthor(criteria string, dateInterval *DateTimeInterval) *FindAuthor {
	return &FindAuthor{Find{Criteria: &Criteria{criteria}, Interval: dateInterval}}
}

func NewFindTitle(criteria string, dateInterval *DateTimeInterval) *FindTitle {
	return &FindTitle{Find{Criteria: &Criteria{criteria}, Interval: dateInterval}}
}
func NewTextSearch(book, criteria string, dateInterval *DateTimeInterval) *TextSearch {
	return &TextSearch{Book: book, Criteria: &Criteria{criteria}, Interval: dateInterval}
}
func NewInterval(from, to time.Time) *DateTimeInterval {
	return &DateTimeInterval{From: from, To: to}
}
func NewTime(date, format string) time.Time {
	t, err := time.Parse(format, date)
	if err != nil {
		panic(err)
	}
	return t
}

func NewString(s interface{}) string {
	return string(s.(*token.Token).Lit)
}

func NewSpace(s interface{}) *Space {
	return &Space{string(s.(*token.Token).Lit)}
}
