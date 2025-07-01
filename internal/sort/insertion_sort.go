package sort

func InsertionSort(data []int) []int {
	for i := range data {
		j := i
		for j > 0 && data[j-1] > data[j] {
			l := j - 1
			r := j

			data[l], data[r] = data[r], data[l]

			j--
		}
	}
	return data
}
