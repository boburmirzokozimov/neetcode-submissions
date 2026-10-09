func sortArray(nums []int) []int {
    if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2

	l := sortArray(nums[:mid])
	r := sortArray(nums[mid:])

	return merge(l, r)
}

func merge(left, right []int)[]int {
	l, r := 0, 0
	result := make([]int, 0, len(left) + len(right))

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}	

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result 
}
