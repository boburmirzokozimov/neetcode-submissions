func minSubArrayLen(target int, nums []int) int {
	best := math.MaxInt

	l := 0
	currSum := 0

	for r := 0; r < len(nums); r++ {
		currSum = currSum + nums[r]

		for currSum >= target {
			best = min(best, r-l+1)
			currSum = currSum - nums[l]
			l++
		}
	}

	if best == math.MaxInt {
		return 0
	}


	return best
}
