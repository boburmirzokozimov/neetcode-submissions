func threeSum(nums []int) [][]int {
	out := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		a := nums[i]

		if a > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums) - 1

		for l < r {
			three := a + nums[l] + nums[r]

			if three > 0 {
				r--
			} else if three < 0 {
				l++
			} else {
				out = append(out, []int{a, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return out
}
