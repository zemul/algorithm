package _sum

import "sort"

/**
https://leetcode.cn/problems/3sum/description/?envType=study-plan-v2&envId=top-interview-150

15. 三数之和 medium

给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。


示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
示例 3：

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。

*/

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 需要和上一次枚举的数不相同
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		j, k := i+1, len(nums)-1
		// 在 i + 1 ... nums.length - 1 中查找相加等于 -nums[i] 的两个数
		for j < k {
			sum := nums[j] + nums[k]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				// 去重
				jVal, kVal := nums[j], nums[k]
				for j < k && nums[j] == jVal {
					j++
				}
				for j < k && nums[k] == kVal {
					k--
				}

			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}
