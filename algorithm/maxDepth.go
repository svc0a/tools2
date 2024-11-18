package algorithm

// 二叉树的最大深度。
// 题目描述
// 给定一个二叉树，找出其最大深度。
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
// 示例:
// 给定二叉树 [3,9,20,null,null,15,7]，
//
//	  3
//	 / \
//	9  20
//	  /  \
//	 15   7
//
// 返回它的最大深度 3 。
// 解题思路
// 这个问题可以通过递归来解决。对于每一个节点，其深度为其左右子树的深度的最大值加1。
func MaxDepth(in []int) int {
	t := buildTree(in)
	return maxDepth(t)
}

func maxDepth(t *node) int {
	if t == nil {
		return 0
	}
	l := maxDepth(t.left)
	r := maxDepth(t.right)
	if l > r {
		return l + 1
	}
	return r + 1
}

func buildTree(in []int) *node {
	if len(in) == 0 {
		return nil
	}
	root := &node{value: -1}
	root.left = buildTree(in[1:])
	root.right = buildTree(in[:])
	return root
}

type node struct {
	value int
	left  *node
	right *node
}
