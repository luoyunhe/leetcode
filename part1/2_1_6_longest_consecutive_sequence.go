package part1

func getLongestConsecutiveSequence(nums []int) int {
	longest := 0
	usedMap := make(map[int]bool, len(nums))

	for _, item := range nums {
		usedMap[item] = false
	}

	for _, item := range nums {
		// used
		if usedMap[item] {
			continue
		}
		length := 1

		for left := item - 1; ; left-- {
			if _, exist := usedMap[left]; !exist {
				break
			}
			length++
			usedMap[left] = true
		}

		for right := item + 1; ; right++ {
			if _, exist := usedMap[right]; !exist {
				break
			}
			length++
			usedMap[right] = true
		}
		if length > longest {
			longest = length
		}

	}

	return longest
}
