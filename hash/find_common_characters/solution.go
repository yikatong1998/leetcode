package find_common_characters

// leetcode 1002 easy

//给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（ 包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答案。
//
//
//示例 1：
//
//输入：words = ["bella","label","roller"]
//输出：["e","l","l"]
//示例 2：
//
//输入：words = ["cool","lock","cook"]
//输出：["c","o"]
//
//
//提示：
//
//1 <= words.length <= 100
//1 <= words[i].length <= 100
//words[i] 由小写英文字母组成

//执行用时：12 ms, 在所有 Go 提交中击败了 7.89% 的用户
//内存消耗：4.3 MB, 在所有 Go 提交中击败了 7.24% 的用户
//通过测试用例：83 / 83

func commonChars(words []string) []string {
	var m = make(map[string]int)
	var s1 = words[0]
	for _, c := range s1 {
		m[string(c)] = m[string(c)] + 1
	}

	if len(words) <= 1 {
		return map2Slice(m)
	}

	for _, s := range words[1:] {
		var tmp = make(map[string]int)
		for _, c := range s {
			if m[string(c)] > 0 {
				m[string(c)]--
				tmp[string(c)] = tmp[string(c)] + 1
			}
		}
		m = tmp
	}

	return map2Slice(m)
}

func map2Slice(m map[string]int) []string {
	var ret = make([]string, 0)
	for s, n := range m {
		for i := 0; i < n; i++ {
			ret = append(ret, s)
		}
	}

	return ret
}
