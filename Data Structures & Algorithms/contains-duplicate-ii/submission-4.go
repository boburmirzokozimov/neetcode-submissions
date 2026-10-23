func containsNearbyDuplicate(nums []int, k int) bool {
	store := make(map[int]bool)
	r := 0
	for i := 0; i < len(nums); i++ {
		if i - r > k {
			delete(store, nums[r])
			r++
		}
		val := nums[i]
		if store[val] {
			return true
		} 
		store[val] = true
	}

	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
