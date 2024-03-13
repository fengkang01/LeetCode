package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	res []int
}

func Constructor(root *TreeNode) FindElements {
	root.Val = 0
	result := &FindElements{
		res: make([]int, 0),
	}
	result.res = append(result.res, root.Val)
	if root.Left != nil {
		coverLeftTree(root.Left, 0, result)
	}
	if root.Right != nil {
		coverRightTree(root.Right, 0, result)
	}
	return *result
}

func coverLeftTree(root *TreeNode, rootVal int, result *FindElements) {
	root.Val = rootVal*2 + 1
	result.res = append(result.res, root.Val)
	if root.Left != nil {
		coverLeftTree(root.Left, root.Val, result)
	}
	if root.Right != nil {
		coverRightTree(root.Right, root.Val, result)
	}
}

func coverRightTree(root *TreeNode, rootVal int, result *FindElements) {
	root.Val = rootVal*2 + 2
	result.res = append(result.res, root.Val)
	if root.Left != nil {
		coverLeftTree(root.Left, root.Val, result)
	}
	if root.Right != nil {
		coverRightTree(root.Right, root.Val, result)
	}
}

func (this *FindElements) Find(target int) bool {
	for _, v := range this.res {
		if v == target {
			return true
		}
	}
	return false
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
