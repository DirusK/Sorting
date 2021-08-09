package sort

func selectSort(arr []int){
	for i:=0;i<len(arr)-1;i++{
		min := i
		for j:=i+1;j<len(arr);j++{
			if arr[min]>arr[j]{
				min = j
			}
		}
		if min!=i{
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
}