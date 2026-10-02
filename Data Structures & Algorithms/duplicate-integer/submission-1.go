func hasDuplicate(nums []int) bool {
    store := make(map[int]bool)
	for _, num := range nums {
		_, ok := store[num]
		if ok {
			return true
		}
		store[num] = true
	}

	return false
}
