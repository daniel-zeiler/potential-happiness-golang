package array_problems

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	var solution [][]int
	for _, interval := range intervals {
		if len(solution) == 0 || solution[len(solution)-1][1] < interval[0] {
			solution = append(solution, interval)
		} else {
			solution[len(solution)-1][1] = max(solution[len(solution)-1][1], interval[1])
		}
	}
	return solution
}

func arrayPairSum(input []int) int {
	sort.Ints(input)
	var result = 0
	for x := 0; x < len(input); x += 2 {
		result += min(input[x], input[x+1])
	}
	return result
}

func SortArrayByParity(input []int) []int {
	var writePointer = 0
	for readPointer := 0; readPointer < len(input); readPointer++ {
		if input[readPointer]%2 == 0 {
			input[writePointer], input[readPointer] = input[readPointer], input[writePointer]
			writePointer++
		}
	}
	return input
}

func replaceElements(input []int) []int {
	var maxSoFar = -1
	for readPointer := len(input) - 1; readPointer != -1; readPointer -= 1 {
		input[readPointer], maxSoFar = maxSoFar, max(maxSoFar, input[readPointer])
	}
	return input
}

func permuteBacktracking(numbersRemaining []int, tempResult []int, res *[][]int) {
	if len(numbersRemaining) == 0 {
		*res = append(*res, tempResult)
	}
	for index, value := range numbersRemaining {
		tempResult = append(tempResult, value)
		permuteBacktracking(append(append([]int{}, numbersRemaining[:index]...), numbersRemaining[index+1:]...), tempResult, res)
		tempResult = tempResult[:len(tempResult)-1]
	}
}

func permute(nums []int) [][]int {
	var res [][]int
	var tempResult []int
	permuteBacktracking(nums, tempResult, &res)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
