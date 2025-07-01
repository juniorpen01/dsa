package sort

func SelectionSort(data []int) []int {
	for i := range data {
		min_idx := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min_idx] {
				min_idx = j
			}
		}
		data[i], data[min_idx] = data[min_idx], data[i]
	}
	return data
}
