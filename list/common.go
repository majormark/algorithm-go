package list

type Node struct {
	Val  int
	Next *Node
}

type DoubleNode struct {
	Val  int
	Next *DoubleNode
	Last *DoubleNode
}
