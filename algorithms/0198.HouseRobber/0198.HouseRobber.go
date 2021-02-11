// Source : https://leetcode.com/problems/house-robber/
// Author : Zhonghuan Gao
// Date   : 2020-02-11

/**********************************************************************************
*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.

Example 2:
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.

Constraints:
0 <= nums.length <= 100
0 <= nums[i] <= 400
*
***********************************************************************************/

// 题意：打家劫舍
package main

/**
解法一：动态规划
设dp[i]为前i个房子能偷的最大金额
状态转移方程：dp[i] = max(dp[i-1], dp[i-2]+nums[i])
时间复杂度：O(n)
空间复杂度：O(n)
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	if n == 1 {
		return dp[0]
	}
	dp[1] = max(nums[0], nums[1])
	if n == 2 {
		return dp[1]
	}
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}