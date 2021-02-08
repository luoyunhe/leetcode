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

func getLongestConsecutiveSequenceV2(nums []int) int {
	clusterMap := make(map[int]int, len(nums))

	longest := 0
	for _, item := range nums {
		if _, exist := clusterMap[item]; exist {
			continue
		}
		// new cluster
		clusterMap[item] = 1

		// right cluster exist
		if _, exist := clusterMap[item+1]; exist {
			length := mergeCluster(clusterMap, item, item+1)
			if length > longest {
				longest = length
			}
		}
		// left cluster exist
		if _, exist := clusterMap[item-1]; exist {
			length := mergeCluster(clusterMap, item-1, item)
			if length > longest {
				longest = length
			}
		}

	}

	return longest
}

func mergeCluster(clusterMap map[int]int, left int, right int) int {
	upper := right + clusterMap[right] - 1
	lower := left - clusterMap[left] + 1
	length := upper - lower + 1
	clusterMap[upper] = length
	clusterMap[lower] = length
	return length
}
