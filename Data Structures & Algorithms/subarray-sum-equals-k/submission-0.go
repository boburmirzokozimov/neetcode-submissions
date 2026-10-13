func subarraySum(nums []int, k int) int {
	store := map[int]int{0: 1}

	res, currSum := 0, 0
	for _, num := range nums {
		currSum += num
		diff := currSum - k
		res += store[diff]
		store[currSum]++
	}

	return res
}
