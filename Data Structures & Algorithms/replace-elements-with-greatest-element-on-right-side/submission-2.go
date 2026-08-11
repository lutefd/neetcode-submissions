func replaceElements(arr []int) []int {
	maxRight := -1

	for i := len(arr) -1; i >= 0; i--{
		curr := arr[i]
		arr[i] = maxRight

		if curr > maxRight {
			maxRight = curr
		}
		
	}
	return arr
}
