package main

import "fmt"

func main() {

	arr1 := []int{5, 4, 3, 2, 1}

	fmt.Println("Unsorted Array ----", arr1)

	sortedArr := SelectionSort(arr1)
	fmt.Println("Sort array by selection sort ----", sortedArr)

	BubbleSort := BubbleSort(arr1)
	fmt.Println("Sort array by bubble sort ----", BubbleSort)
}
