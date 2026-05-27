package node

const (
	PageSize = 4096
	MaxPages = 1024
)

var (
	leaf         LeafNode
	internalNode InternalNode

	CommonHeaderSize       = leaf.CommonHeader.Size()
	InternalNodeHeaderSize = CommonHeaderSize + internalNode.Header.Size()
	InternalNodeSize       = internalNode.Size()
	InternalNodeCellSize   = internalNode.ICells[0].Size()
	InternalNodeMaxCells   = uint32(len(internalNode.ICells))

	LeafNodeHeaderSize = CommonHeaderSize + leaf.Header.Size()
	LeafNodeSize       = leaf.Size()
	LeafNodeCellSize   = leaf.Cells[0].Size()
	LeafNodeMaxCells   = uint32(len(leaf.Cells))

	RightSplitCount = (LeafNodeMaxCells + 1) / 2
	LeftSplitCount  = LeafNodeMaxCells + 1 - RightSplitCount
)

// FindChildByKey returns the index of the child which should contain
//
//	the given key.
func (d *InternalNode) FindChildByKey(key uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func (d *InternalNode) Child(childIdx uint32) (ptr *uint32) { _ = "STUB: not implemented"; return nil }
