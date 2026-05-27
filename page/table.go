package page

import (
	"github.com/auxten/go-sqldb/node"
)

type Table struct {
	Pager       *Pager
	RootPageIdx uint32
}

// Seek the page of key, if not exist then return the place key should be
// for the later INSERT.
func (table *Table) Seek(key uint32) (cursor *Cursor, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (table *Table) Insert(row *node.Row) (err error) { _ = "STUB: not implemented"; return nil }

// Must be leaf node

func (table *Table) leafNodeSeek(pageIdx uint32, key uint32) (cursor *Cursor, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Walk the btree

func (table *Table) internalNodeSeek(pageIdx uint32, key uint32) (cursor *Cursor, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (table *Table) CreateNewRoot(rightChildPageIdx uint32) (err error) {
	_ = "STUB: not implemented"
	/*
	  Handle splitting the root.
	  Old root copied to new page, becomes left child.
	  Address of right child passed in.
	  Re-initialize root page to contain the new root node.
	  New root node points to two children.
	*/return nil
}

// copy whatever kind of node to leftChildPage, and set nonRoot

// 重新初始化 root page，root page 将会有一个 key，两个子节点

func (table *Table) InternalNodeInsert(parentPageIdx uint32, childPageIdx uint32) (err error) {
	_ = "STUB: not implemented"
	/*
	  Add a new child/key pair to parent that corresponds to child
	*/return nil
}

/* Replace right child */

/* Make room for the new cell */

func (table *Table) Select() { _ = "STUB: not implemented"; return }

func (table *Table) Prepare() { _ = "STUB: not implemented"; return }
