package main

/*
https://leetcode.cn/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150

189. 轮转数组 medium

*给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

提示：

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105
*/

/*
* 找规律
该方法基于如下的事实：当我们将数组的元素向右移动 k 次后，尾部 kmodn 个元素会移动至数组头部，其余元素向后移动 kmodn 个位置。

该方法为数组的翻转：我们可以先将所有元素翻转，这样尾部的 kmodn 个元素就被移至数组头部，然后我们再翻转 [0,kmodn−1] 区间的元素和 [kmodn,n−1] 区间的元素即能得到最后的答案。

操作	结果
原始数组						1 2 3 4 5 6 7
翻转所有元素					7 6 5 4 3 2 1
翻转 [0,kmodn−1] 区间的元素	5 6 7 4 3 2 1
翻转 [kmodn,n−1] 区间的元素	5 6 7 1 2 3 4
*/
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
	return
}

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		t := nums[i]
		nums[i] = nums[j]
		nums[j] = t
		i++
		j--
	}
	return
}

// 新数组遍历赋值法, 空间O(n)
func rotate2(nums []int, k int) {
	length := len(nums)
	res := make([]int, length)

	for i := range nums {
		res[(i+k)%length] = nums[i]
	}
	copy(nums, res)
	return
}
