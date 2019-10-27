package miscellaneous

// findSecondLargest 找到二叉查找树的第二大节点，
// 并返回其值
func findSecondLargest(root *TreeNode) int {

	if root == nil || (root.Left == nil && root.Right == nil) {
		return -1
	}

	current := root
	for current != nil {
		// 找到最大值的节点后，判断最大值节点情况
		// 只有两种情况，一种是有左子树，那么左子树的最大值就是
		// 左子树的最大值
		// 另外一种情况是，没有左右子树，那么当前节点的父节点
		// 是次大值
		// 没有第三种情况
		if current.Left != nil && current.Right == nil {
			return findSecondLargest(current.Left)
		}
		if current.Right != nil && current.Right.Left == nil && current.Right.Right == nil {
			return current.Val
		}

		current = current.Right
	}

	return -1
}

func findLargest(root *TreeNode) int {
	current := root
	for current != nil {
		if current.Right == nil {
			return current.Val
		}
		current = current.Right
	}
	return -1
}
