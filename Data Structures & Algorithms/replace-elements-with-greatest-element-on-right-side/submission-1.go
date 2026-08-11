func replaceElements(arr []int) []int {
	pointer := len(arr) -2
	lastGreatestSeen := arr[len(arr) - 1]
	arr[len(arr)-1] = -1
	
	for pointer >= 0 {
		if arr[pointer] > lastGreatestSeen{
			lastGreatestSeen, arr[pointer] = arr[pointer], lastGreatestSeen
			
		} else {
			arr[pointer] = lastGreatestSeen
		}
		
		pointer--
	}
	return arr
}
