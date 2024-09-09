package product_of_array_except_self

/*
*
https://leetcode.cn/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-interview-150

238. 除自身以外数组的乘积 medium

给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。

示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]

提示：

2 <= nums.length <= 105
-30 <= nums[i] <= 30
保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内

进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）

*/

/**
思考

题目禁止使用除法，如果允许，可能陷入一个陷阱
正常思路，可以直接求整个数组的乘积，再除以当前值，但是如果数组中含有0的话，会导致结果错误
*/

// 左右乘积列表
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	// left[i] 为索引 i 左侧所有元素的乘积
	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	right[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		ans[i] = left[i] * right[i]
	}

	return ans
}
