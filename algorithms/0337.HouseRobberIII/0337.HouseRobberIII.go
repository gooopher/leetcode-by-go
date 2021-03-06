// Source : https://leetcode.com/problems/house-robber-iii/
// Author : Zhonghuan Gao
// Date   : 2020-03-06

/**********************************************************************************
*
The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.
Determine the maximum amount of money the thief can rob tonight without alerting the police.

Example 1:
Input: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

Output: 7
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.

Example 2:
Input: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
*
***********************************************************************************/

// 题意：打家劫舍III
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
解法一：递归
思路：maxRes = max(root.Val + root.Left.Left.Val + root.Left.Right.Val + root.Right.Left.Val + root.Right.Right.Val,
					root.Left.Val + root.Right.Val)
测试用例：[1,2,3,5,4,8,9,null,null,6,7]
 */
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nextLevel := rob(root.Left) + rob(root.Right)

	nextNextLevel := 0
	if root.Left != nil {
		nextNextLevel += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		nextNextLevel += rob(root.Right.Left) + rob(root.Right.Right)
	}

	return max(root.Val + nextNextLevel, nextLevel)
}

/**
解法二：动态规划，参考198和213，是其的扩展
思路：
	由题可知当前节点的子树上被选择的节点最大权值和有两种可能，一种是包含当前节点，即当前节点被选中；一种是不包含，即当前节点不被选中。
	设 l[0], l[1] 分别为当前节点的左子节点被选时和不被选时最大权值和， r[0], r[1] 分别为当前节点的右子节点被选时和不被选时最大权值和
	则可能当前节点被选和不被选时的最大权值和分别为 {node.value + l[1] + r[1], max(l[0], l[1], r[0], r[1])}
时间复杂度：O(n)
空间复杂度：O(n)
参考：https://leetcode-cn.com/problems/house-robber-iii/solution/da-jia-jie-she-iii-by-leetcode-solution/
*/
func rob2(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}
