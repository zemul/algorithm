package candy

/**
135. 分发糖果 困难 https://leetcode-cn.com/problems/candy/
n个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。



示例 1：

输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
示例 2：

输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。

*/
// 两次遍历
// https://www.bilibili.com/video/BV1cS4y1z7so/?spm_id_from=333.337.search-card.all.click&vd_source=fd8414f99601e4f73736044cdca613aa
func candy(ratings []int) int {
	var ans int
	var left = make([]int, len(ratings))
	var right = make([]int, len(ratings))
	for i, v := range ratings {
		if i > 0 && v > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}
	for i := range ratings {
		ans += max(left[i], right[i])
	}
	return ans
}
