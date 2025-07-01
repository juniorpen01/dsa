package sort

func BubbleSort(data []int) {
	swapping := true
	end := len(data)
	for swapping {
		swapping = false
		for i := 1; i < end; i++ {
			if l, r := i-1, i; data[l] > data[r] {
				data[l], data[r] = data[r], data[l]
				swapping = true
			}
		}
		end--
	}
}
