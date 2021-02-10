package part1

import (
	"math"
	"sort"
)

func closestThreeSum(nums []int, target int) (closest int) {
	minGap := math.MaxInt64

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			gap := abs(sum, target)
			if gap < minGap {
				minGap = gap
			}
			if sum > target {
				k--
			} else if sum < target {
				j++
			} else {
				return 0
			}

		}

	}

	return minGap
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}

}
