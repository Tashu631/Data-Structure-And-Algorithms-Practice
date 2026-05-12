package main

import (
	"fmt"
)

func PalindromeNumber() {
	num := 120
	temp := num
	reverseNumber := 0
	for num > 0 {
		rem := num % 10
		reverseNumber = reverseNumber*10 + rem
		num = num / 10
	}
	if reverseNumber == temp {
		fmt.Println(true)
	} else {
		fmt.Print(false)
	}
	// fmt.Println(reverseNumber)
}
