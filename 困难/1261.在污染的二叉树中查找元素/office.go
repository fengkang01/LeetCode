package main

type FindElements2 struct {
	res map[int]bool
}

func Constructor2(root *TreeNode) FindElements2 {
	root.Val = 0
	result := &FindElements2{
		res: make(map[int]bool),
	}
	result.res[root.Val] = true
	dfs(root, 0, result)
	return *result
}

func dfs(root *TreeNode, rootVal int, result *FindElements2) {
	if root == nil {
		return
	}
	result.res[rootVal] = true
	dfs(root.Left, root.Val*2+1, result)
	dfs(root.Right, root.Val*2+2, result)
}

func (this *FindElements2) Find(target int) bool {
	_, ok := this.res[target]
	return ok
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
