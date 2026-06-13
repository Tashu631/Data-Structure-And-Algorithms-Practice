package main

func MergeSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	MergeSort(arr, low, mid)
	MergeSort(arr, mid+1, high)
	Merge(low, high, mid, arr)
}
func Merge(low, high, mid int, arr []int) {
	i := low
	j := mid + 1
	var temp []int
	for i <= mid && j <= high {
		if arr[i] < arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}
	for i <= mid {
		temp = append(temp, arr[i])
		i++
	}

	for j <= high {
		temp = append(temp, arr[j])
		j++
	}

	for k := 0; k < len(temp); k++ {
		arr[k+low] = temp[k]
	}

}
