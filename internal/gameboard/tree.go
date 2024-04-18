package gameboard

import "fmt"

type Node struct {
	Data   *Board
	Depth  int
	Childs []*Node
	Move   int
}

func NewNode(b *Board, depth int, move int) *Node {
	return &Node{
		Data:  b,
		Depth: depth,
		Move:  move,
	}
}

func (n *Node) AddChild(child *Node) {
	n.Childs = append(n.Childs, child)
}

func (n *Node) String() string {
	return fmt.Sprintf("Depth : %d\n"+n.Data.String(), n.Depth)
}
