package main

func InsertionSort(arr []int)[]int{
   for i :=1;i<len(arr);i++{
		key := arr[i]
		j:= i-1
		for j >=0 && arr[j] > key{
			 arr[j]=arr[j+1]
			 j--
		}
		arr[j+1]=key
	 }

	return arr
}