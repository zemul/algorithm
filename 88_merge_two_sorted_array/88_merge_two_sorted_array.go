package main

import "sort"

/**
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 指向 nums1 中最后一个元素的位置
	curr := m + n - 1
	// 指向 nums1 和 nums2 中剩余元素的索引
	i, j := m-1, n-1

	// 遍历两个数组，直到一个数组为空
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[curr] = nums1[i]
			i--
		} else {
			nums1[curr] = nums2[j]
			j--
		}
		curr--
	}
	// 如果 nums2 还有剩余元素，将它们复制到 nums1 的前面
	for j >= 0 {
		nums1[curr] = nums2[j]
		j--
		curr--
	}
}

// 1. 最直观的方法是先将数组 nums 2放进数组 nums 1 的尾部，然后直接对整个数组进行排序。
func merge1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
