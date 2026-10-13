func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	store := make(map[int]bool)

	for _, num := range nums {
		store[num] = true
	}

	longest := 1
	for i := 0; i < len(nums); i++ {
		currLongest := 1
		next := nums[i] + 1

		if store[nums[i] - 1] {
			continue
		}
		for store[next] {
			currLongest++
			next = next + 1
		}
		longest = max(longest, currLongest)
	}
//[2,20,4,10,3,4,5]

	return longest
}
