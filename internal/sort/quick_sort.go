package sort

func QuickSort(data []int) []int {
	quickSort(data, 0, len(data)-1)
	return data
}

func quickSort(data []int, low, high int) {
	if low < high {
		middle := partition(data, low, high)
		quickSort(data, low, middle-1)
		quickSort(data, middle+1, high)
	}
}

func partition(data []int, low, high int) int {
	pivot := data[high]
	i := low
	for j := low; j < high; j++ {
		if data[j] < pivot {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[high] = data[high], data[i]
	return i
}
