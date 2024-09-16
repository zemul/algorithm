package two_sum

/**
https://leetcode.cn/problems/two-sum/description/?envType=study-plan-v2&envId=top-interview-150

1. 两数之和 easy

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。



示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]

*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for idx, num := range nums {
		if tarIdx, has := m[target-num]; has {
			return []int{idx, tarIdx}
		}
		m[num] = idx
	}

	return []int{}
}

/** 错误例子
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for idx, num:=range nums{
		m[num] = idx
	}

	//【 target = 6 num = 3，会返回相同坐标 】
	for idx, num :=range nums{
		if tarIdx, has := m[target-num];has{
			return []int{idx,tarIdx}
		}
	}
	return []int{}

}
*/
