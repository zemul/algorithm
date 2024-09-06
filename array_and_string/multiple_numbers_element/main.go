package main

/*
*
https://leetcode.cn/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150

169. 多数元素

给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：

输入：nums = [3,2,3]
输出：3
示例 2：

输入：nums = [2,2,1,1,1,2,2]
输出：2

提示：
n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109

进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
*/

// 摩尔投票法 空间O(1)
// 若记 众数 的票数为 +1 ，非众数 的票数为 −1 ，则一定有所有数字的 票数和 >0,因为众数的票数大于其余所有数字的票数和。
// 利用此特性，遇到相同的数则票数 +1 ，遇到不同的数则票数 -1 ，通过不断抵消不同数的票数，最终获得票数最多的数。
func majorityElement2(nums []int) int {
	var candidate, votes int
	for _, num := range nums {
		if votes == 0 {
			candidate = num
		}
		if num == candidate {
			votes++
		} else {
			votes--
		}
	}
	return candidate
}

// 哈希表 空间O(n)
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] > len(nums)/2 {
			return num
		}
	}
	return 0
}
