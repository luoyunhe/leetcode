package part1

import (
	"fmt"
	"sort"
)

func genKey(res []int) string {
	return fmt.Sprintf("%d-%d-%d-%d", res[0], res[1], res[2], res[3])
}

func fourSum(nums []int, target int) (resultList [][]int) {
	if len(nums) < 4 {
		return
	}
	sort.Ints(nums)

	set := make(map[string]struct{})

	for a := 0; a < len(nums)-3; a++ {
		for b := a + 1; b < len(nums)-2; b++ {
			c := b + 1
			d := len(nums) - 1
			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum > target {
					d--
				} else if sum < target {
					c++
				} else {
					result := []int{nums[a], nums[b], nums[c], nums[d]}
					if _, exist := set[genKey(result)]; exist {
						d--
						c++
					} else {
						resultList = append(resultList, result)
						set[genKey(result)] = struct{}{}
					}
				}
			}
		}

	}

	return
}
