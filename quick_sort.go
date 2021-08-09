package sort

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	split := partition(arr)

	quickSort(arr[:split])
	quickSort(arr[split:])
}

func partition(arr []int) int {
	left := 0
	right := len(arr) - 1

	pivot := arr[len(arr)/2]

	for {
		for ; arr[left] < pivot; left++ {
		}
		for ; arr[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		arr[left], arr[right] = arr[right], arr[left]
	}
}

