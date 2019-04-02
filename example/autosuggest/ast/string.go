package ast

import (
	"fmt"
)

func (fa *FindAuthor) String() string {
	return fmt.Sprintf("Find Author - Criteria: %s, Interval: %s", fa.Criteria, fa.Interval)
}

func (fa *FindTitle) String() string {
	return fmt.Sprintf("Find Title - Criteria: %s, Interval: %s", fa.Criteria, fa.Interval)
}

func (ts *TextSearch) String() string {
	return fmt.Sprintf("TextSearch - Book: %s, Criteria: %s, Interval: %s", ts.Book, ts.Criteria, ts.Interval)
}

func (c *Criteria) String() string {
	return fmt.Sprintf("%s", c.S)
}

func (dt *DateTimeInterval) String() string {
	return fmt.Sprintf("%s -- %s", dt.From, dt.To)
}
