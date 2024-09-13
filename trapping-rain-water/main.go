package trapping_rain_water

/**
42. 接雨水 困难

https://leetcode.cn/problems/trapping-rain-water/description/?envType=study-plan-v2&envId=top-interview-150

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：


输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

*/
// todo 单调栈 双指针

// 动态规划,记录每个元素分别向左和向右扫描并记录左边和右边的最大高度，然后计算每个下标位置能接的雨水量。
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	left := make([]int, n)
	right := make([]int, n)
	left[0] = height[0]
	right[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i])
	}

	for j := n - 2; j >= 0; j-- {
		right[j] = max(right[j+1], height[j])
	}

	var ans int
	for i, h := range height {
		ans += min(left[i], right[i]) - h
	}
	return ans
}
