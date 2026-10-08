func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	
    m := len(nums) / 2

	l := sortArray(nums[:m])
	r := sortArray(nums[m:])

	return merge(l, r)
}

func merge(left []int, right []int)[]int {
	result := make([]int, 0, len(left) + len(right))

	i, j := 0, 0 

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
