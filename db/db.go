package db

import (
	"github.com/auxten/go-sqldb/page"
)

func Open(fileName string) (t *page.Table, err error) { _ = "STUB: not implemented"; return nil, nil }

// New database file, initialize page 0 as leaf node.

func Close(t *page.Table) (err error) { _ = "STUB: not implemented"; return nil }
