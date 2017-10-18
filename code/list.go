package main

import (
	"fmt"
)

type ListNode struct {
	data interface{}
	next *ListNode
}

func (list *ListNode) Init() {
	list.data = 0
	list.next = nil
}

func (list *ListNode) Add(prev, next, node *ListNode) {
	/*
		    if nil == prev {
				node.next = next
				list = node //错误写法
				return
			}
	*/
	prev.next = node
	node.next = next
	return
}

func (list *ListNode) AddHeadNext(node *ListNode) {
	list.Add(list, list.next, node)
}

func (list *ListNode) Delete(prev, next *ListNode) {
	prev.next = next
}

func (list *ListNode) DeleteNode(node *ListNode) {
	for pnode := list; pnode != nil; pnode = pnode.next {
		if pnode.next == node {
			list.Delete(pnode, pnode.next)
			break
		}
	}
}

func (list *ListNode) Traverse() {
	for pnode := list; nil != pnode; pnode = pnode.next {
		fmt.Println(pnode)
	}
}

var listnodes = []ListNode{{1, nil}, {2, nil}}

func main() {
	fmt.Println("list test")
	list := ListNode{0, nil}

	//for _, node := range listnodes{} //错误写法，node只是listnodes[i]的拷贝
	for i, _ := range listnodes {
		list.AddHeadNext(&listnodes[i])
	}

	list.Traverse()
}
