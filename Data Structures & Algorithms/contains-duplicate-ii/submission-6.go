func containsNearbyDuplicate(nums []int, k int) bool {
	store := make(map[int]bool)

	l := 0
	for r := 0; r < len(nums); r++ {
		
		if r - l > k {
			delete(store, nums[l])
			l++
		} 

		val := nums[r]

		if store[val] {
			return true
		}

		store[val] = true
	}

	return false
}

