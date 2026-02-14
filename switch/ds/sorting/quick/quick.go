package main 
func quickSort(arr []int, low int, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)

		quickSort(arr, low, pivotIndex-1)   // left side
		quickSort(arr, pivotIndex+1, high)  // right side
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]  // choose last element as pivot
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}