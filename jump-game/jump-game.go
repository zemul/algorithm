package main

/**

55. 跳跃游戏  medium
https://leetcode.cn/problems/jump-game/description/?envType=study-plan-v2&envId=top-interview-150

给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。

提示：
1 <= nums.length <= 104
0 <= nums[i] <= 105
*/

func canJump(nums []int) bool {
	maxIdx := 0
	for idx, num := range nums {
		// 当前位置不可达
		if maxIdx < idx {
			return false
		}
		maxIdx = max(maxIdx, idx+num)
	}
	return true
}

/*
*
dp[i]= max(i+nums[i],dp[i-1])
条件： dp[i]可达 i <= dp[i-1]
*/
func canJump2(nums []int) bool {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		// 如果当前位置不可达，注意退出，继续循环会导致从一个不可达的位置开始寻找最大值导致永远返回true
		if i > dp[i-1] {
			return false
		} else {
			dp[i] = max(dp[i-1], i+nums[i])
		}
	}
	return dp[len(nums)-1] >= len(nums)-1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
