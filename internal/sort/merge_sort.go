package sort

func MergeSort(data []int) []int {
	length := len(data)

	if length < 2 {
		return data
	}

	l := MergeSort(data[:length/2])
	r := MergeSort(data[length/2:])

	return merge(l, r)
}

func merge(l, r []int) []int {
	final := []int{}

	var i, j int

	for i < len(l) && j < len(r) {
		if l, r := l[i], r[j]; l <= r {
			final = append(final, l)
			i++
		} else {
			final = append(final, r)
			j++
		}
	}

	for i < len(l) {
		final = append(final, l[i])
		i++
	}

	for j < len(r) {
		final = append(final, r[j])
		j++
	}

	return final
}
