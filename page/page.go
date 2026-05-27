package page

import (
	"os"

	"github.com/auxten/go-sqldb/node"
)

type Page struct {
	// Either InternalNode or LeafNode
	InternalNode *node.InternalNode
	LeafNode     *node.LeafNode
}

func (p *Page) GetMaxKey() uint32 { _ = "STUB: not implemented"; return 0 }

type Pager struct {
	File    *os.File
	fileLen int64
	PageNum uint32  // PageNum is the boundary of db memory page.
	Pages   []*Page // Page pointer slice, nil member indicates cache missing.
}

func PagerOpen(fileName string) (pager *Pager, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get file length

// dbFile length must be n * node.PageSize, node.PageSize is usually 4096

func (p *Pager) GetPage(pageIdx uint32) (page *Page, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cache miss
// If pageIdx within data file, just read,
// else just return blank page which will be flushed to db file later.

// Load page from file

// Empty new page will be leaf node

// Leaf node

// Internal node

func (p *Pager) Flush(pageIdx uint32) (err error) { _ = "STUB: not implemented"; return nil }
