package part1

func twoSum(nums []int, target int) (index1, index2 int) {
	indexMap := make(map[int]int, len(nums))
	for index, item := range nums {
		indexMap[item] = index
	}
	for index, item := range nums {
		if idx, exist := indexMap[target-item]; exist {
			return index + 1, idx + 1
		}
	}
	return
}
