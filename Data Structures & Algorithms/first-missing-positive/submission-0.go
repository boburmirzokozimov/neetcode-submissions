func firstMissingPositive(nums []int) int {
	n := len(nums)

	// Put every number x in position x-1
	for i := 0; i < n; i++ {
		for nums[i] >= 1 &&
			nums[i] <= n &&
			nums[i] != nums[nums[i]-1] {

			correctIndex := nums[i] - 1
			nums[i], nums[correctIndex] = nums[correctIndex], nums[i]
		}
	}

	// Find the first number that isn't in its correct position
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 1..n are all present
	return n + 1
}