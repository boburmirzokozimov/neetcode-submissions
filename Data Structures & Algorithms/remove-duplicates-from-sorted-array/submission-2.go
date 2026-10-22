func removeDuplicates(nums []int) int {
	writeIndex := 1

	for read := 1; read < len(nums); read++ {
		if nums[read] != nums[read-1] {
			nums[writeIndex] = nums[read]
			writeIndex++
		}
	}

	//[1,1,2,3,4]
	//[2,10,10,30,30,30]
	return writeIndex
}
