func merge(nums1 []int, m int, nums2 []int, n int) {
    last := m + n - 1

    // Merge in reverse order
    for m > 0 && n > 0 {
        if nums1[m-1] > nums2[n-1] {
            nums1[last] = nums1[m-1]
            m--
        } else {
            nums1[last] = nums2[n-1]
            n--
        }
        last--
    }

    // Fill nums1 with leftover nums2 elements
    for n > 0 {
        nums1[last] = nums2[n-1]
        n--
        last--
    }
}