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

/* EXPLANATION:
Reversing a Number - Iterative Approach:

Goal: Reverse the digits of a number
Example: 123456789 → 987654321

Algorithm:
1. Initialize reversed = 0
2. Extract last digit using modulo: digit = num % 10
3. Add extracted digit to reversed: reversed = reversed*10 + digit
4. Remove last digit from num: num = num / 10
5. Repeat until num becomes 0

Step-by-step for 123:
- rem = 123 % 10 = 3, reversed = 0*10 + 3 = 3, num = 123/10 = 12
- rem = 12 % 10 = 2, reversed = 3*10 + 2 = 32, num = 12/10 = 1
- rem = 1 % 10 = 1, reversed = 32*10 + 1 = 321, num = 1/10 = 0
- Loop ends, result = 321

For 123456789:
- Output: 987654321

Time Complexity: O(number of digits) = O(log n)
Space Complexity: O(1)
*/
