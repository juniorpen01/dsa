package algorithms

func PowerSet(input []int) [][]int {
	if len(input) == 0 {
		return nil
	}
	all_subsets := [][]int{{}}
	for _, e := range input {
		subsets_with_element := [][]int{}
		for _, subset := range all_subsets {
			new_subset := append(subset, e)
			subsets_with_element = append(subsets_with_element, new_subset)
		}
		all_subsets = append(all_subsets, subsets_with_element...)
	}
	return all_subsets
}
