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

type cacheItem struct {
	cIndex int
	dIndex int
}

func fourSum2(nums []int, target int) (resultList [][]int) {
	if len(nums) < 4 {
		return
	}
	sort.Ints(nums)

	cache := make(map[int][]*cacheItem)

	for c := 2; c < len(nums)-1; c++ {
		for d := c + 1; d < len(nums); d++ {
			cache[nums[c]+nums[d]] = append(cache[nums[c]+nums[d]], &cacheItem{
				cIndex: c,
				dIndex: d,
			})
		}
	}
	set := make(map[string]struct{})

	for a := 0; a < len(nums)-3; a++ {
		for b := a + 1; b < len(nums)-2; b++ {
			list := cache[target-nums[a]-nums[b]]
			for _, item := range list {
				if item.cIndex <= b {
					continue
				}
				result := []int{nums[a], nums[b], nums[item.cIndex], nums[item.dIndex]}
				if _, exist := set[genKey(result)]; exist {
					continue
				} else {
					resultList = append(resultList, result)
					set[genKey(result)] = struct{}{}
				}

			}

		}

	}
	return
}
