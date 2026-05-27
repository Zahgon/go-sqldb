package planner

import (
	"github.com/auxten/go-sqldb/parser"
)

func (plan *Plan) Insert(ast *parser.InsertTree) (count int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// 这里暂时都假定我们插入的 Schema 是固定的 node.Row 类型
// 根据 InsertTree.Columns 的字段顺序，我们强制类型转换还原出 node.Row 结构
