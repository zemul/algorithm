package container_with_most_water

/*
*
https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-interview-150

11. 盛最多水的容器 medium

给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

示例 1：

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1
*/
func maxArea(height []int) int {
	var ans int
	var l, r = 0, len(height) - 1
	// 如果移动长的一边,宽已经由短的一边固定了,长度在不断缩减,所以以后的面积只会比现在小；
	// 如果移动短的一边，即使长度变短了,但是宽度和长度的乘积可能会变大；最后就是将最大值返回

	for l < r {
		val := min(height[l], height[r]) * (r - l)
		ans = max(ans, val)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return ans

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
