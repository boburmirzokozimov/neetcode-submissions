func minSubArrayLen(target int, nums []int) int {
	l := 0
	minL := math.MaxInt
	totalSum := 0

	for r := 0; r < len(nums); r++ {
		totalSum += nums[r]

		for totalSum >= target {
			minL = min(minL, r-l+1)
			totalSum -= nums[l]
			l++
		}
	}

	if minL == math.MaxInt {
		return 0
	}

	return minL
}
