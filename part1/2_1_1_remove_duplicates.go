package part1

// remove duplicates from sorted array
func removeDuplicates1(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
		}
	}
	return nums[:index+1]
}
