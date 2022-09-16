package intersection_of_two_arrays

// leetcode 349 easy

// 给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
//
//
//
//示例 1：
//
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
//示例 2：
//
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//解释：[4,9] 也是可通过的
//
//
//提示：
//
//1 <= nums1.length, nums2.length <= 1000
//0 <= nums1[i], nums2[i] <= 1000

// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗： 3 MB, 在所有 Go 提交中击败了 25.55% 的用户
//通过测试用例： 55 / 55

func intersection(nums1 []int, nums2 []int) []int {
	var hash = make(map[int]int)
	for _, num := range nums1 {
		hash[num] = 1
	}

	var ret = make([]int, 0)
	for _, num := range nums2 {
		if hash[num] > 0 {
			hash[num]--
			ret = append(ret, num)
		}
	}

	return ret
}
