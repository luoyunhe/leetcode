package part1

func removeElement(nums []int, target int) (newNums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != target {
			nums[index] = nums[i]
			index++
		}
	}
	newNums = nums[:index]
	return
}
