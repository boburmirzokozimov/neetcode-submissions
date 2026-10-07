func majorityElement(nums []int) int {
    maxEl := nums[0]
	maxElCount := 1

	for i := 1; i < len(nums); i++ {
		if maxEl == nums[i] {
			maxElCount++
		} else {
			maxElCount--
		}
		if maxElCount == 0 {
			maxElCount = 1
			maxEl = nums[i]
		} 
	}

	return maxEl
}
//[5,5,1,1,1,5,5]
