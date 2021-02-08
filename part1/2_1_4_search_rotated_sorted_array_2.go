package part1

func searchRotatedSortedArray2(nums []int, target int) int {
	first := 0
	last := len(nums)
	for first != last {
		mid := first + (last-first)/2
		if nums[mid] == target {
			return mid
		}
		if nums[first] < nums[mid] {
			if nums[first] <= target && target < nums[mid] {
				last = mid
			} else {
				first = mid + 1
			}
		} else if nums[first] > nums[mid] {
			if nums[mid] < target && target <= nums[last-1] {
				first = mid + 1
			} else {
				last = mid
			}
		} else {
			first++
		}
	}
	return -1
}
