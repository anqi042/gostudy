package main

import (
	"fmt"
)

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) init(val interface{}) {
	t.Val = val
	t.Left = nil
	t.Right = nil
}

func (t *TreeNode) compare(value1, value2 interface{}) bool {
	switch value1.(type) {
	case int:
		if val1, ok := value1.(int); ok {
			if val2, ok := value2.(int); ok {
				return val1 >= val2
			}
		}
	}

	return false
}

func (t *TreeNode) insert(value interface{}) {
	if t.compare(value, t.Val) {
		if nil == t.Right {
			t.Right = &TreeNode{value, nil, nil}
		} else {
			t.Right.insert(value)
		}
	} else {
		if nil == t.Left {
			t.Left = &TreeNode{value, nil, nil}
		} else {
			t.Left.insert(value)
		}
	}
}

/* 中序遍历 */
func (t *TreeNode) traverse() {
	if nil != t {
		fmt.Println(t.Val)
		t.Left.traverse()
		t.Right.traverse()
	}
}

func main() {
	fmt.Println("------------------tree pratice-----------------")
	a := []int{10, 7, 3, 1, 8, 13, 11, 12}

	tree := new(TreeNode)
	tree.init(a[0])

	for _, val := range a[1:] {
		tree.insert(val)
	}

	tree.traverse()
}
