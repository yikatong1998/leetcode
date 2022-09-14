package spiral_matrix_ii

// leetcode 59 medium

// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

//执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗： 1.9 MB, 在所有 Go 提交中击败了 98.44% 的用户
//通过测试用例： 20 / 20

var ret [][]int
var n, num, i, j int

func generateMatrix(m int) [][]int {
	ret = nil
	n, num, i, j = m, 1, 0, 0
	for k := 0; k < n; k++ {
		ret = append(ret, make([]int, n, n))
	}

	for {
		fill()
		if n <= 0 {
			break
		}
	}

	return ret
}

func fill() {
	ret[i][j] = num

	for k := 0; k < n-1; k++ {
		j++
		num++
		ret[i][j] = num
	}

	for k := 0; k < n-1; k++ {
		i++
		num++
		ret[i][j] = num
	}

	for k := 0; k < n-1; k++ {
		j--
		num++
		ret[i][j] = num
	}

	for k := 0; k < n-2; k++ {
		i--
		num++
		ret[i][j] = num
	}

	n -= 2
	num += 1
	i = i
	j += 1
}
