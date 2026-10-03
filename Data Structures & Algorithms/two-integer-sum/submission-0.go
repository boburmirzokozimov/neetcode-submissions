func twoSum(nums []int, target int) []int {
    store := make(map[int]int)
	//7-3 = 4/0
	//7-4 = 3/1
	for idx, num := range nums {
		pos, exists := store[num]
		if exists {
			return []int{pos, idx}
		}
		val := target - num
		store[val] = idx
	}

	return []int{0, 0}
}
