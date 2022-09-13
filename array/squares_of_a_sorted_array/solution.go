package squares_of_a_sorted_array

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//
//
//示例 1：
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
//示例 2：
//
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]
//
//
//提示：
//
//1 <= nums.length <= 104
//-104 <= nums[i] <= 104
//nums 已按 非递减顺序 排序
//
//
//进阶：
//
//请你设计时间复杂度为 O(n) 的算法解决本问题

//执行用时： 20 ms, 在所有 Go 提交中击败了 97.07% 的用户
//内存消耗： 6.7 MB, 在所有 Go 提交中击败了 9.44% 的用户
//通过测试用例： 137 / 137

func sortedSquares(nums []int) []int {
	var ret []int

	var i, j = -1, 0
	for ; j < len(nums); j++ {
		if nums[j] >= 0 {
			i = j - 1
			break
		}
	}
	if j == len(nums) {
		i = len(nums) - 1
	}

	for i > -1 || j < len(nums) {
		if j >= len(nums) {
			ret = append(ret, nums[i]*nums[i])
			i--
			continue
		}

		if i <= -1 {
			ret = append(ret, nums[j]*nums[j])
			j++
			continue
		}

		if -nums[i] <= nums[j] {
			ret = append(ret, nums[i]*nums[i])
			i--
			continue
		}

		if -nums[i] > nums[j] {
			ret = append(ret, nums[j]*nums[j])
			j++
			continue
		}
	}

	return ret
}
