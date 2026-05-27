package parser

import (
	"text/scanner"
)

type StatementType string

// SQL type tokens
const (
	UNSUPPORTED = "N/A"
	SELECT      = "SELECT"
	FROM        = "FROM"
	WHERE       = "WHERE"
	LIMIT       = "LIMIT"
	INSERT      = "INSERT"
	INTO        = "INTO"
	VALUES      = "VALUES"
	ASTERISK    = "*"
)

type Parser struct {
	s scanner.Scanner
}

type SelectTree struct {
	Projects []string
	Table    string
	Where    []string
	Limit    int64
}

type InsertTree struct {
	Table   string
	Columns []string
	Values  [][]string
}

func (p *Parser) GetSQLType(sql string) StatementType {
	_ = "STUB: not implemented"
	return *new(StatementType)
}

/*
ParseSelect is a simple select statement parser.
It's just a demo of SELECT statement parser skeleton.
Currently, the most complex SQL supported here is something like:

	SELECT * FROM foo WHERE id < 3 LIMIT 1;

Even SQL-92 standard is far more complex.
For a production ready SQL parser, see: https://github.com/auxten/postgresql-parser
*/
func (p *Parser) ParseSelect(sel string) (ast *SelectTree, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//log.Print(txt)

// token FROM is scanned, try to get the table name here
// FROM ?

// if projects are all constant value, source table is not necessary.
// eg.  SELECT 1;

// WHERE

// WHERE is not necessary

// token WHERE is scanned, try to get the WHERE clause.

// token LIMIT is scanned, try to get the limit

/*
ParseInsert can parse a simple INSERT statement, eg.

	 	INSERT INTO table_name VALUES (value1, value2, …)
		or
		INSERT INTO table_name(column1, column2, …) VALUES (value1, value2, …)
*/
func (p *Parser) ParseInsert(insert string) (ast *InsertTree, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Table name

// try get colNames

//log.Print(txt)

// VALUES has been scanned try to get (value1, value2), (value3, value4)

// next row

//log.Print(txt)

// Check if column count identical
