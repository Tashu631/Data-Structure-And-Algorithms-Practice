package main

func SelectionSort(arr []int) []int {
	// 5 4 3 2 1
	for i := 0; i < len(arr); i++ { // i =0
		// ele := arr[i]                        // ele = 5
		minEle := i
		for j := i + 1; j < len(arr); j++ { // j =1  2   3  4   5 condition false
			if arr[j] < arr[minEle] {
				minEle = j 
			}
		}
		
		arr[i],arr[minEle] = arr[minEle],arr[i] 
	}

	return arr
}
