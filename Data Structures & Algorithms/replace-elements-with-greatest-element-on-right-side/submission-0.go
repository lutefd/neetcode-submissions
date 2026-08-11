func replaceElements(arr []int) []int {
	pointer := len(arr) -1
	var lastGreatestSeen int

	for pointer >= 0 {
	
		if arr[pointer] > lastGreatestSeen{
			lastGreatestSeen, arr[pointer] = arr[pointer], lastGreatestSeen
			
		} else {
			arr[pointer] = lastGreatestSeen
		}
			if pointer == len(arr) -1 {
			arr[pointer] = -1
		}
		pointer--
	}
	return arr
}
