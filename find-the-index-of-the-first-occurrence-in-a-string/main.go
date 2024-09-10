package main

/**
https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/?envType=study-plan-v2&envId=top-interview-150

easy

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。

示例 1：

输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
示例 2：

输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
*/

// kmp 时间复杂度：O(n+m) 空间O(m)
// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/solutions/575568/shua-chuan-lc-shuang-bai-po-su-jie-fa-km-tb86/?envType=study-plan-v2&envId=top-interview-150
// todo
func strStr(haystack string, needle string) int {
	return -1
}

// 暴力匹配 时间复杂度：O(n×m)，n 是字符串 haystack 的长度，m 是字符串 needle 的长度。
func strStr1(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	for i := 0; i <= hlen-nlen; i++ {
		for j := 0; j < nlen; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == nlen-1 {
				return i
			}
		}
	}
	return -1
}

// 错误代码，
// 部分匹配失败，将 nidx（needle 的索引）直接重置为 0，而没有考虑已经匹配到的部分。
//
//	mississippi issip
func strStr2(haystack string, needle string) int {
	var nidx = 0
	for hidx := 0; hidx < len(haystack); hidx++ {
		if haystack[hidx] != needle[nidx] {
			nidx = 0
			continue
		} else {
			nidx++
		}
		if nidx == len(needle) {
			return hidx - len(needle) + 1
		}

	}
	return -1
}
