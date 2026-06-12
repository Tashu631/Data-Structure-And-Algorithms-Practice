package main

func BubbleSort(nums []int)[]int{
		for i := range nums{
			for j := i+1;j<len(nums)-i-1;j++{
				  if nums[i]>nums[j]{
						nums[i],nums[j]=nums[j],nums[i]
					}
			}
		}
		return nums
}