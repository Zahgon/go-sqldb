package planner

import (
	"github.com/auxten/go-sqldb/node"
	"github.com/auxten/go-sqldb/page"
	"github.com/auxten/go-sqldb/parser"
)

func (plan *Plan) SelectPrepare(ast *parser.SelectTree) (filteredPipe chan *node.Row, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the start of table

// If the key not found in the table and the page to insert that key
// has 0 cells. We got the end of the table.

/*
	The code below demonstrates a simple "Volcano Model" query plan.
	 For more please refer to https://doi.org/10.1109/69.273032
*/

// Fetch rows from storage pages

// Filter rows according the ast.Where

// Count row count for LIMIT clause.

func (plan *Plan) fetchRow(table *page.Table) (row *node.Row, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Move cursor to next leaf

// 已经移动到了最右的的叶子节点

func isRowFiltered(where []string, row *node.Row) (filtered bool, err error) {
	_ = "STUB: not implemented"
	// This is a very dirty hack to use Eval to evaluate the Where statement.
	return false, nil
}

/*
	type Row struct {
		Id       uint32
		Sex      byte
		Age      uint8
		Username [32]byte
		Email    [128]byte
		Phone    [64]byte
	}
*/
