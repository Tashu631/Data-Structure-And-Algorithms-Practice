package main

import "fmt"

func main() {
	arr := []int{-1, 0, 1, 3, 5, 9, 12}

	target := 9

	result := Search(arr, target)
	fmt.Println(result)
}
