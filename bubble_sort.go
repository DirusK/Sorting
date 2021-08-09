package sort

func bubbleSort(arr []int){
	for i := 0; i < len(arr); i++ {
		for j := len(arr)-1; j > i ; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}		
		}
	}
}
