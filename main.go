package sort

import (
	"fmt"
	"time"
	"math/rand"
)

func test(sort func([]int), array []int){
	start := time.Now()
	sort(array)
	elapsed := time.Since(start)
	fmt.Println("Time required: ", elapsed.Milliseconds())
}

func generate(array []int){
	for i:=0;i<len(array);i++{
		array[i]=rand.Intn(11)-10
	}
}
func main() {
	array := make([]int,10)
	generate(array)
	fmt.Println("Initial array: ", array)
	test(bubbleSort,array)
	fmt.Println("Bubble Sort: ",array)
	fmt.Println("***************************************************")


	generate(array)
	fmt.Println("Initial array: ", array)
	test(quickSort,array)
	fmt.Println("Quick Sort: ",array)
	fmt.Println("***************************************************")
}
