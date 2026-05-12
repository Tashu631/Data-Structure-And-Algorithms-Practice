package main

import (
	"fmt"
)

func ArmstrongNumber() {
	num := 371
	temp := num
	sum := 0
	for num > 0 {
		rem := num % 10
		sum = sum + (rem * rem * rem)
		num = num / 10
	}

	if sum == temp {
		fmt.Println(true)
	} else {
		fmt.Print(false)
	}
}
