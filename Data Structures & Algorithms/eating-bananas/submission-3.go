func minEatingSpeed(piles []int, h int) int {
    l, r := 1, 0

    for _, p := range piles {
        r = max(r, p)
    }

    for l < r {
        mid := l + (r-l)/2
        totalSpent := 0

        for _, p := range piles {
            totalSpent += (p + mid - 1) / mid
        }

        if totalSpent <= h {
            r = mid
        } else {
            l = mid + 1
        }
    }

    return l
}