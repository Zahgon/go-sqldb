package planner

import (
	"github.com/auxten/go-sqldb/node"
	"github.com/auxten/go-sqldb/page"
)

type Plan struct {
	table          *page.Table
	cursor         *page.Cursor
	UnFilteredPipe chan *node.Row
	FilteredPipe   chan *node.Row
	LimitedPipe    chan *node.Row
	ErrorsPipe     chan error
	Stop           chan bool
}

func NewPlan(t *page.Table) (p *Plan) { _ = "STUB: not implemented"; return nil }
