package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, nil, nil)
}

func isBST(node *TreeNode, min *string, max *string) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Data <= *min) || (max != nil && node.Data >= *max) {
		return false
	}

	if !isBST(node.Left, min, &node.Data) {
		return false
	}

	if !isBST(node.Right, &node.Data, max) {
		return false
	}

	return true
}
