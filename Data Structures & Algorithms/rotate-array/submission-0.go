func rotate(nums []int, k int) {
	k = k % len(nums)
	n := len(nums) - 1

	reverse(nums, 0, n)
	reverse(nums, 0, k-1)
	reverse(nums, k, n)
	//[1,2,3,4,5,6,7,8]
	//[8,7,1,2,3,4,5,6]
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
