package main

/**
easy

https://leetcode.cn/problems/length-of-last-word/description/?envType=study-plan-v2&envId=top-interview-150

给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大
子字符串
。

示例 1：

输入：s = "Hello World"
输出：5
解释：最后一个单词是“World”，长度为 5。
示例 2：

输入：s = "   fly me   to   the moon  "
输出：4
解释：最后一个单词是“moon”，长度为 4。
示例 3：

输入：s = "luffy is still joyboy"
输出：6
解释：最后一个单词是长度为 6 的“joyboy”。

提示：

1 <= s.length <= 104
s 仅有英文字母和空格 ' ' 组成
s 中至少存在一个单词
*/

func lengthOfLastWord(s string) int {
	var (
		idx int = len(s) - 1
		ans int
	)
	for s[idx] == ' ' {
		idx--
	}

	for idx >= 0 && s[idx] != ' ' {
		ans++
		idx--
	}
	return ans
}

// self
func lengthOfLastWord2(s string) int {
	end := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			end = i
			break
		}
	}
	for i := end; i >= 0; i-- {
		if s[i] == ' ' {
			return end - i
		}
	}

	// +1，因为 len(s[:end]) 计算的长度不包括 end 索引处的字符
	return len(s[:end]) + 1
}
