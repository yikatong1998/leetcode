package valid_anagram

import "reflect"

// leetcode 242 easy

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//
//
//
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//示例 2:
//
//输入: s = "rat", t = "car"
//输出: false
//
//
//提示:
//
//1 <= s.length, t.length <= 5 * 104
//s 和 t 仅包含小写字母
//
//
//进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

//执行用时： 16 ms, 在所有 Go 提交中击败了 10.75% 的用户
//内存消耗： 3.6 MB, 在所有 Go 提交中击败了 5.19% 的用户
//通过测试用例： 37 / 37

func isAnagram(s string, t string) bool {
	var record1 = record(s)
	var record2 = record(t)

	return reflect.DeepEqual(record1, record2)
}

func record(s string) map[string]int {
	var record = make(map[string]int)
	for _, c := range s {
		record[string(c)] = record[string(c)] + 1
	}
	return record
}
