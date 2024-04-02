// https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/description/?envType=daily-question&envId=2024-01-28
package solution

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	count := 0

	for xIndex := range matrix {
		for yIndex := range matrix[xIndex] {
			count += countSubmatrixFrom(xIndex, yIndex, target, matrix)
		}
	}

	return count
}

func countSubmatrixFrom(x, y, target int, matrix [][]int) int {
	count := 0

	// dp := make([][]int)

	for dx := 0; dx < len(matrix)-x; dx++ {
		for dy := 0; dy < len(matrix[dx])-y; dy++ {
			sum := 0
			for cx := x; cx <= x+dx; cx++ {
				subSlice := matrix[cx][y : y+dy+1]
				for _, value := range subSlice {
					sum += value
				}
			}
			if sum == target {
				count++
			}
		}
	}
	return count
}
