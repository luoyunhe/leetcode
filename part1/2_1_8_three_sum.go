package part1

import "sort"

func threeSum(nums []int, target int) (result [][]int) {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			tmp := nums[i] + nums[j] + nums[k]
			if tmp > target {
				k--
				for nums[k] == nums[k+1] && j < k {
					k--
				}
			} else if tmp < target {
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
				j++
				for nums[k] == nums[k+1] && j < k {
					k--
				}
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			}

		}

	}

	return result
}
