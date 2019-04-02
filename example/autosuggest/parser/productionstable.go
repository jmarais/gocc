// Code generated by gocc; DO NOT EDIT.

package parser

import (
    "time"
	"github.com/goccmack/gocc/example/autosuggest/ast"
)

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Start	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Start : FindAuthor	<< &ast.Search{FindAuthor: X[0].(*ast.FindAuthor)} , nil >>`,
		Id:         "Start",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Search{FindAuthor: X[0].(*ast.FindAuthor)} , nil
		},
	},
	ProdTabEntry{
		String: `Start : FindTitle	<< &ast.Search{FindTitle: X[0].(*ast.FindTitle)} , nil >>`,
		Id:         "Start",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Search{FindTitle: X[0].(*ast.FindTitle)} , nil
		},
	},
	ProdTabEntry{
		String: `Start : TextSearch	<< &ast.Search{TextSearch: X[0].(*ast.TextSearch)} , nil >>`,
		Id:         "Start",
		NTType:     1,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &ast.Search{TextSearch: X[0].(*ast.TextSearch)} , nil
		},
	},
	ProdTabEntry{
		String: `FindAuthor : "find" AuthorKeyWord Space "with" Space id Space "at" Space TimeRange	<< ast.NewFindAuthor(ast.NewString(X[5]), X[9].(*ast.DateTimeInterval)) , nil >>`,
		Id:         "FindAuthor",
		NTType:     2,
		Index:      4,
		NumSymbols: 10,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFindAuthor(ast.NewString(X[5]), X[9].(*ast.DateTimeInterval)) , nil
		},
	},
	ProdTabEntry{
		String: `FindTitle : "find" TitleKeyWord Space "with" Space id Space "at" Space TimeRange	<< ast.NewFindTitle(ast.NewString(X[5]), X[9].(*ast.DateTimeInterval)) , nil >>`,
		Id:         "FindTitle",
		NTType:     3,
		Index:      5,
		NumSymbols: 10,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFindTitle(ast.NewString(X[5]), X[9].(*ast.DateTimeInterval)) , nil
		},
	},
	ProdTabEntry{
		String: `AuthorKeyWord : "author"	<<  >>`,
		Id:         "AuthorKeyWord",
		NTType:     4,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AuthorKeyWord : Space "author"	<<  >>`,
		Id:         "AuthorKeyWord",
		NTType:     4,
		Index:      7,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TitleKeyWord : "title"	<<  >>`,
		Id:         "TitleKeyWord",
		NTType:     5,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TitleKeyWord : Space "title"	<<  >>`,
		Id:         "TitleKeyWord",
		NTType:     5,
		Index:      9,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TextSearch : "textsearch" Space id Space "with" Space id Space "at" Space TimeRange	<< ast.NewTextSearch(ast.NewString(X[2]), ast.NewString(X[6]), X[10].(*ast.DateTimeInterval)) , nil >>`,
		Id:         "TextSearch",
		NTType:     6,
		Index:      10,
		NumSymbols: 11,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTextSearch(ast.NewString(X[2]), ast.NewString(X[6]), X[10].(*ast.DateTimeInterval)) , nil
		},
	},
	ProdTabEntry{
		String: `TextSearch : "textsearch" Space id Space "with" Space id	<< ast.NewTextSearch(ast.NewString(X[2]), ast.NewString(X[6]), nil) , nil >>`,
		Id:         "TextSearch",
		NTType:     6,
		Index:      11,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTextSearch(ast.NewString(X[2]), ast.NewString(X[6]), nil) , nil
		},
	},
	ProdTabEntry{
		String: `TimeRange : TimeRangeKeyWord DateTimeInterval	<< X[1], nil >>`,
		Id:         "TimeRange",
		NTType:     7,
		Index:      12,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `TimeRangeKeyWord : "timerange"	<<  >>`,
		Id:         "TimeRangeKeyWord",
		NTType:     8,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TimeRangeKeyWord : Space "timerange"	<<  >>`,
		Id:         "TimeRangeKeyWord",
		NTType:     8,
		Index:      14,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DateTimeInterval : DateTime "--" DateTime	<< ast.NewInterval(X[0].(time.Time), X[2].(time.Time)), nil >>`,
		Id:         "DateTimeInterval",
		NTType:     9,
		Index:      15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInterval(X[0].(time.Time), X[2].(time.Time)), nil
		},
	},
	ProdTabEntry{
		String: `DateTime : Date Space TimeHMS	<< ast.NewTime(X[0].(string) + " " + X[2].(string), "2006-01-02 15:04:05"), nil >>`,
		Id:         "DateTime",
		NTType:     10,
		Index:      16,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTime(X[0].(string) + " " + X[2].(string), "2006-01-02 15:04:05"), nil
		},
	},
	ProdTabEntry{
		String: `DateTime : Date Space TimeHMS Space	<< ast.NewTime(X[0].(string) + " " + X[2].(string), "2006-01-02 15:04:05"), nil >>`,
		Id:         "DateTime",
		NTType:     10,
		Index:      17,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTime(X[0].(string) + " " + X[2].(string), "2006-01-02 15:04:05"), nil
		},
	},
	ProdTabEntry{
		String: `Date : Number Number Number	<< X[0].(string)+X[1].(string)+X[2].(string), nil >>`,
		Id:         "Date",
		NTType:     11,
		Index:      18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(string)+X[1].(string)+X[2].(string), nil
		},
	},
	ProdTabEntry{
		String: `Date : Space Number Number Number	<< X[1].(string)+X[2].(string)+X[3].(string), nil >>`,
		Id:         "Date",
		NTType:     11,
		Index:      19,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1].(string)+X[2].(string)+X[3].(string), nil
		},
	},
	ProdTabEntry{
		String: `TimeHMS : Number ":" Number ":" Number	<< X[0].(string)+ast.NewString(X[1])+X[2].(string)+ast.NewString(X[3])+X[4].(string), nil >>`,
		Id:         "TimeHMS",
		NTType:     12,
		Index:      20,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(string)+ast.NewString(X[1])+X[2].(string)+ast.NewString(X[3])+X[4].(string), nil
		},
	},
	ProdTabEntry{
		String: `Number : int_lit	<< ast.NewString(X[0]), nil >>`,
		Id:         "Number",
		NTType:     13,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewString(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Number : dayMonth	<< ast.NewString(X[0]), nil >>`,
		Id:         "Number",
		NTType:     13,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewString(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Number : number	<< ast.NewString(X[0]), nil >>`,
		Id:         "Number",
		NTType:     13,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewString(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Space : space	<< ast.NewSpace(X[0]), nil >>`,
		Id:         "Space",
		NTType:     14,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSpace(X[0]), nil
		},
	},
}
