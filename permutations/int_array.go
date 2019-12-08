package permutations

func GenerateInt(a []int) [][]int {
	return generate(a, len(a))
}

func generate(a []int, k int) [][]int {
	if k == 1 {
		tmp := make([]int, len(a))
		copy(tmp, a)
		return [][]int{tmp}
	}

	res := generate(a, k - 1)

	for i := 0; i < k-1; i++ {
		if k % 2 == 0 {
			tmp := a[i]
			a[i] = a[k-1]
			a[k-1] = tmp
		} else {
			tmp := a[0]
			a[0] = a[k-1]
			a[k-1] = tmp
		}

		for _, item := range generate(a, k-1) {
			res = append(res, item)
		}
	}

	return res
}
