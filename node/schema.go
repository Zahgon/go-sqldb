package node

import (
	"io"
)

var (
	RowSize = (&Row{}).Size()
)

/*
Id       uint32
Sex      byte
Age      uint8
Username [32]byte
Email    [128]byte
Phone    [64]byte
*/
func PrintRow(row *Row) { _ = "STUB: not implemented"; return }

func WriteRow(w io.Writer, row *Row) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func dumpConst() { _ = "STUB: not implemented"; return }
