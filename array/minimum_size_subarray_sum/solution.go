package minimum_size_subarray_sum

// leetcode 209 medium

//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
//
//
//示例 1：
//
//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//示例 2：
//
//输入：target = 4, nums = [1,4,4]
//输出：1
//示例 3：
//
//输入：target = 11, nums = [1,1,1,1,1,1,1,1]
//输出：0
//
//
//提示：
//
//1 <= target <= 109
//1 <= nums.length <= 105
//1 <= nums[i] <= 105
//
//
//进阶：
//
//如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。

//执行用时： 24 ms, 在所有 Go 提交中击败了 88.48% 的用户
//内存消耗： 8.1 MB, 在所有 Go 提交中击败了 88.55% 的用户
//通过测试用例： 20 / 20

func minSubArrayLen(target int, nums []int) int {
	var i, j, sum int
	var ret = len(nums) + 1

	for ; j < len(nums); j++ {
		sum += nums[j]
		if sum < target {
			continue
		} else {
			width := j - i + 1
			ret = min(ret, width)
			for ; i <= j; i++ {
				sum -= nums[i]
				if sum < target {
					i++
					break
				} else {
					width := j - i
					ret = min(ret, width)
				}
			}
		}
	}

	if ret == len(nums)+1 {
		return 0
	} else {
		return ret
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
