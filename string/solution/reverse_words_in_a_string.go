package solution

// leetcode 151 medium

// 给你一个字符串 s ，请你反转字符串中 单词 的顺序。
//
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
//
// 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
//
// 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
//
//
//
// 示例 1：
//
// 输入：s = "the sky is blue"
// 输出："blue is sky the"
// 示例 2：
//
// 输入：s = "  hello world  "
// 输出："world hello"
// 解释：反转后的字符串中不能存在前导空格和尾随空格。
// 示例 3：
//
// 输入：s = "a good   example"
// 输出："example good a"
// 解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
//
// 提示：
//
// 1 <= s.length <= 104
// s 包含英文大小写字母、数字和空格 ' '
// s 中 至少存在一个 单词
//
// 进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用 O(1) 额外空间复杂度的 原地 解法。

// 执行用时： 4 ms, 在所有 Go 提交中击败了 33.38% 的用户
//内存消耗： 2.7 MB, 在所有 Go 提交中击败了 55.94% 的用户
//通过测试用例： 58 / 58

func reverseWords(s string) string {
	ret := make([]byte, 0)
	j := len(s) - 1
	for j >= 0 {
		for j >= 0 && string(s[j]) == " " {
			j--
		}
		i := j - 1
		for i >= 0 && string(s[i]) != " " {
			i--
		}
		if i+1 >= 0 {
			ret = append(ret, []byte(s[i+1:j+1])...)
			ret = append(ret, byte(' '))
		}

		j = i - 1
	}

	if len(ret) > 0 {
		return string(ret[:len(ret)-1])
	} else {
		return string(ret)
	}
}
