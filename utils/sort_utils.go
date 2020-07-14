package utils

// BubbleSort will sort the given elements in ascending order
func BubbleSort(elements []int) []int {
	sorted := true
	for sorted {
		sorted = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				sorted = true
			}
		}
	}
	return elements
}
