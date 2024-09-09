package main

/**
45. 跳跃游戏 II  medium

https://leetcode.cn/problems/jump-game-ii/?envType=study-plan-v2&envId=top-interview-150

给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

0 <= j <= nums[i]
i + j < n
返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。


示例 1:
输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

示例 2:
输入: nums = [2,3,0,1,4]
输出: 2

提示:
1 <= nums.length <= 104
0 <= nums[i] <= 1000
题目保证可以到达 nums[n-1]
*/

// 我的解
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	steps := 0
	maxPoint := 0
	end := 0
	for i := 0; i < len(nums); i++ {
		maxPoint = max(i+nums[i], maxPoint)
		if i == end {
			steps++
			end = maxPoint
		}
		if end >= len(nums)-1 {
			return steps
		}
	}
	return steps

}

// 官方解
// 维护当前能够到达的最大下标位置，记为边界。从左到右遍历数组，到达边界时，更新边界并将跳跃次数增加 1。
func jump2(nums []int) int {
	steps := 0    // 到达maxPoint 需要的步数
	maxPoint := 0 // 下一跳最大可达位置
	end := 0

	// i < len(nums)-1,注意最后一个元素不遍历，因为steps记录的是到达maxPoint需要的步数
	// 到达倒数第二个元素时，如果到达边界，steps++并更新下一步的maxPoint，否则本step可以直接到达最后一个元素
	for i := 0; i < len(nums)-1; i++ {
		maxPoint = max(maxPoint, i+nums[i])
		if i == end {
			steps++
			end = maxPoint
		}
	}
	return steps
}
