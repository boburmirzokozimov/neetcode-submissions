func subarraySum(nums []int, k int) int {
	seen := map[int]int{0: 1} // empty prefix: lets subarrays starting at index 0 count
	sum, count := 0, 0
	for _, n := range nums {
		sum += n
		count += seen[sum-k] // earlier prefixes that make a subarray summing to k
		seen[sum]++          // record after counting, so a subarray is never empty
	}

	return count
}