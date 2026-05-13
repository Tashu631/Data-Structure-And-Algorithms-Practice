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

/* EXPLANATION:
Palindrome Number Check:

Palindrome: A number that reads the same forwards and backwards
Examples: 121, 1331, 12321 are palindromes
Non-palindromes: 120, 123, 1234

Algorithm:
1. Store original number in 'temp'
2. Extract each digit from right to left using modulo (%)
3. Build reversed number by: reversed = reversed*10 + last_digit
4. Compare reversed number with original
5. Return true if equal, false otherwise

Step-by-step for 121:
- Extract 1: reversed = 0*10 + 1 = 1
- Extract 2: reversed = 1*10 + 2 = 12
- Extract 1: reversed = 12*10 + 1 = 121
- 121 == 121? Yes, so it's palindrome ✓

For 120: reversed = 21, which ≠ 120, so it's NOT palindrome

Output: For 120, prints false
*/
