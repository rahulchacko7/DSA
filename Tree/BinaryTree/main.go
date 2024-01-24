package main

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (t *TreeNode) Insert(value int) *TreeNode {
	if t == nil {
		return &TreeNode{Value: value}
	}
	if value <= t.Value {
		t.Left = t.Left.Insert(value)
	} else {
		t.Right = t.Right.Insert(value)
	}
	return t
}

func (t *TreeNode) Delete(value int) *TreeNode {
	if t == nil {
		return nil
	}
}

func (t *TreeNode) DisplayIn() {
	if t != nil {
		t.Left.DisplayIn()
		fmt.Println(t.Value)
		t.Right.DisplayIn()
	}
}

func main() {
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(88)
	fmt.Println("In-Order Traversal...")
	root.DisplayIn()
	fmt.Println()

}
