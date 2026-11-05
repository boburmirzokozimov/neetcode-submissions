func shipWithinDays(weights []int, days int) int {
	l, r := 1, 0

	for _, w := range weights {
		l = max(l, w)
		r += w
	}
	res := r

	canShip := func(cap int) bool {
		ships, currCap := 1, cap
		for _, w := range weights {
			if currCap - w < 0 {
				ships++
				if ships > days {
					return false
				}
				currCap = cap
			}
			currCap -= w
		}

		return true
	}

	for l <= r {
		cap := l + (r-l)/2
		if canShip(cap) {
			res = min(cap, res)

			r = cap - 1
		} else {
			l = cap + 1
		}
	}

	return res
}
