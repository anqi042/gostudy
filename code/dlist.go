package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

func (list *Node) Init() {
	list.data = 0
	list.next = list
	list.prev = list
}

func (list *Node) _add_(next, prev *Node, node *Node) {
	prev.next = node
	node.prev = prev
	node.next = next
	next.prev = node
}

func (list *Node) AddHead(node *Node) {
	list._add_(list.next, list, node)
}

func (lsit *Node) _delete_(next, prev *Node, node *Node) {
	prev.next = next
	next.prev = prev
}

func (list *Node) Delete(node *Node) {
	list._delete_(node.next, node.prev, node)
}

func (list *Node) Traverse() {
	fmt.Println("list traverse:")
	l := list.next
	for ; l != list; l = l.next {
		fmt.Println(l.data)
	}
}

func main() {
	fmt.Println("dlist practice")
	list := Node{0, nil, nil}
	list.Init()

	node1 := &Node{3, nil, nil}
	list.AddHead(node1)

	node2 := &Node{2, nil, nil}
	list.AddHead(node2)

	list.Traverse()

	list.Delete(node1)

	list.Traverse()
}
