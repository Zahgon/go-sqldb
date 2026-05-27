package page

import (
	"github.com/auxten/go-sqldb/node"
)

type Cursor struct {
	Table      *Table
	PageIdx    uint32
	CellIdx    uint32
	EndOfTable bool
}

func (cursor *Cursor) LeafNodeInsert(key uint32, row *node.Row) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Split leaf node

// Need make room for new cell

func (cursor *Cursor) LeafNodeSplitInsert(key uint32, row *node.Row) (err error) {
	_ = "STUB: not implemented"
	/*
	  Create a new node and move half the cells over.
	  Insert the new value in one of the two nodes.
	  Update parent or create a new parent.
	*/return nil
}

// put new page in the end
// TODO: Page recycle

/*
  All existing keys plus new key should should be divided
  evenly between old (left) and new (right) nodes.
  Starting from the right, move each key to correct position.
*/

/* Update cell count on both leaf nodes */

// parent page is an internal node

func saveToCell(cell *node.Cell, key uint32, row *node.Row) (err error) {
	_ = "STUB: not implemented"
	return nil
}
