package part1

func findKth(a []int, b []int, k int) int {
	la := len(a)
	lb := len(b)
	if la > lb {
		a, b = b, a
	}
	if la == 0 {
		return b[k-1]
	}
	if k == 1 {
		if a[0] > b[0] {
			return b[0]
		} else {
			return a[0]
		}
	}

	da := k / 2
	if da > la {
		da = la
	}

	db := k - da

	if a[da-1] > b[db-1] {
		return findKth(a[:], b[db:], k-db)
	} else if a[da-1] < b[db-1] {
		return findKth(a[da:], b[:], k-da)
	} else {
		return a[da-1]
	}
}
