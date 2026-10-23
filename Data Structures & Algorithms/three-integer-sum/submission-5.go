func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	out := [][]int{}
	//[-1,-1,0,1,2,]
	for i := 0; i < n; i++ {
		curr := nums[i]

		if curr > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			target := -curr

			sum := nums[l] + nums[r]

			if sum == target {
				out = append(out, []int{curr, nums[l], nums[r]})
				l++ 
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if sum > target {
				r--
			} else {
				l++
			}
		}
	}

	return out
}
