package main

import (
	"fmt"
)

func ReverseNumber() {
	num := 123456789
	reverseNumber := 0
	for num > 0 {
		rem := num % 10
		reverseNumber = reverseNumber*10 + rem
		num = num / 10
	}
	fmt.Println(reverseNumber)
}
