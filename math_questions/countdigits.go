package main

import (
	"fmt"
	"math"
)

func CountDigits() {
	num := 12345678910
	count := int(math.Log10(float64(num))) + 1
	fmt.Println("count digits:", count)
	// Process file content
}

/* EXPLANATION:
Counting Digits in a Number using Mathematical Approach:

Method: Using Logarithm (Base 10)
Formula: count = floor(log10(num)) + 1

Why this works:
- log10(100) = 2, floor(2) + 1 = 3 digits ✓
- log10(1000) = 3, floor(3) + 1 = 4 digits ✓
- log10(12345678910) = 10.09, floor(10.09) + 1 = 11 digits

Advantage: Very efficient O(1) time complexity
Alternative: Could use loop and division by 10 (O(log n) time)

Example: For number 12345678910, output will be "count digits: 11"
*/
