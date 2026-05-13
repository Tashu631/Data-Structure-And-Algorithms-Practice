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

/* EXPLANATION:
Armstrong Number (also called Narcissistic Number):
A number is Armstrong if the sum of cubes of its digits equals the number itself.
Example: 371 is Armstrong because 3³ + 7³ + 1³ = 27 + 343 + 1 = 371

Algorithm:
1. Store original number in 'temp' for later comparison
2. Extract each digit using modulo operator (%)
3. Cube each digit and add to sum
4. Remove last digit by dividing by 10
5. Compare final sum with original number
6. Return true if equal, false otherwise

Output: Prints true if 371 is Armstrong, false otherwise
*/
