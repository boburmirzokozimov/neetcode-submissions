func removeDuplicates(nums []int) int {
	writeIndex := 1

	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[r-1] {
			nums[writeIndex] = nums[r]
			writeIndex++
		}
	}
	
	
	//[2,10,10,30,30,30] => 
	return writeIndex
}	
