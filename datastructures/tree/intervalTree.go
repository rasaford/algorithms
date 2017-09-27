package tree

import "github.com/rasaford/algorithms/internal/helper"
import "math"

type iTree struct {
	*rbTree
}

type IntervalNode interface {
	RedBlackNode
	Low() int
	setLow(int)
	High() int
	setHigh(int)
	Max() int
	setMax(int)
}

type iNode struct {
	RedBlackNode
	left, right, parent IntervalNode
	// low & high ends of the intervalls
	low, high int
	// maximum high value of the subtree rooted at this node
	max int
}

func (n *iNode) Low() int      { return n.low }
func (n *iNode) setLow(l int)  { n.low = l }
func (n *iNode) High() int     { return n.high }
func (n *iNode) setHigh(h int) { n.high = h }
func (n *iNode) Max() int      { return n.max }
func (n *iNode) setMax(m int)  { n.max = m }

func NewIntervalTree() *iTree {
	s := &iNode{
		RedBlackNode: &rbNode{
			key:   math.MinInt32,
			color: false,
		},
	}
	return &iTree{
		rbTree: &rbTree{
			root:     s,
			sentinel: s,
		},
	}
}

func (t *iTree) Insert(low, high int, data interface{}) {
	node := &iNode{
		RedBlackNode: &rbNode{
			key:   low,
			value: data,
			color: true,
			left:  t.sentinel,
			right: t.sentinel,
		},
		low:  low,
		high: high,
	}
	t.insert(node)
	node.max = helper.Max(node.high,
		node.left.Max(),
		node.right.Max())
}

func (t *iTree) Delete(low, high int) {
	node := t.Search(low, high)
	t.delete(node)
}

func (t *iTree) Search(low, high int) IntervalNode {
	x := t.root
	for x != t.sentinel && overlap(low, high,
		x.(IntervalNode).Low(),
		x.(IntervalNode).High()) {
		if x.Left() != t.sentinel && x.Left().(IntervalNode).Max() >= low {

		}
	}
	return nil
}

func overlap(aLow, aHigh, bLow, bHigh int) bool {
	return aLow <= bHigh && bLow <= aHigh ||
		bLow <= aHigh && aLow <= bHigh
}
