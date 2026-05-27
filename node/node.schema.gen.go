package node

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type Header struct {
	IsInternal bool
	IsRoot     bool
	Parent     uint32
}

func (d *Header) Size() (s uint64) { _ = "STUB: not implemented"; return 0 }

func (d *Header) Marshal(buf []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *Header) Unmarshal(buf []byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

type InternalNodeHeader struct {
	KeysNum    uint32
	RightChild uint32
}

func (d *InternalNodeHeader) Size() (s uint64) { _ = "STUB: not implemented"; return 0 }

func (d *InternalNodeHeader) Marshal(buf []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *InternalNodeHeader) Unmarshal(buf []byte) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type LeafNodeHeader struct {
	Cells    uint32
	NextLeaf uint32
}

func (d *LeafNodeHeader) Size() (s uint64) { _ = "STUB: not implemented"; return 0 }

func (d *LeafNodeHeader) Marshal(buf []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *LeafNodeHeader) Unmarshal(buf []byte) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type ICell struct {
	Key   uint32
	Child uint32
}

func (d *ICell) Size() (s uint64) { _ = "STUB: not implemented"; return 0 }

func (d *ICell) Marshal(buf []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *ICell) Unmarshal(buf []byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

type InternalNode struct {
	CommonHeader Header
	Header       InternalNodeHeader
	ICells       [510]ICell
}

func (d *InternalNode) Size() (s uint64) { _ = "STUB: not implemented"; return 0 }

// make compiler happy in case k is unused

func (d *InternalNode) Marshal(buf []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *InternalNode) Unmarshal(buf []byte) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type Cell struct {
	Key   uint32
	Value [230]byte
}

func (d *Cell) Size() (s uint64) { _ = "STUB: not implemented"; return 0 }

func (d *Cell) Marshal(buf []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *Cell) Unmarshal(buf []byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

type LeafNode struct {
	CommonHeader Header
	Header       LeafNodeHeader
	Cells        [17]Cell
}

func (d *LeafNode) Size() (s uint64) { _ = "STUB: not implemented"; return 0 }

// make compiler happy in case k is unused

func (d *LeafNode) Marshal(buf []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *LeafNode) Unmarshal(buf []byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

type Row struct {
	Id       uint32
	Sex      byte
	Age      uint8
	Username [32]byte
	Email    [128]byte
	Phone    [64]byte
}

func (d *Row) Size() (s uint64) { _ = "STUB: not implemented"; return 0 }

func (d *Row) Marshal(buf []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *Row) Unmarshal(buf []byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
