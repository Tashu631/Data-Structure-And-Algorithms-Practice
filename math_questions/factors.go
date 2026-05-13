package main

import (
	"fmt"
	"math"
	"sort"
)

func Factors() {
	num := 20
	var divisor []int

	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			divisor = append(divisor, i)
		}
		if (num / i) != i {
			divisor = append(divisor, num/i)
		}
	}
	sort.Ints(divisor)
	fmt.Println(divisor)
}

/* EXPLANATION:
Finding All Factors of a Number:

What are factors? Numbers that divide the given number completely (remainder = 0)
Example: Factors of 20 are 1, 2, 4, 5, 10, 20

Optimized Algorithm:
1. Loop only up to sqrt(num) instead of num (much faster)
2. For each divisor i found, also get its complement (num/i)
3. Avoid adding same number twice if i == num/i (for perfect squares)
4. Sort the divisors in ascending order

Why sqrt optimization works?
- Factors come in pairs: if i divides num, then (num/i) also divides num
- One factor from each pair is always ≤ sqrt(num)
- Time Complexity: O(√n) instead of O(n)

Example: For 20, output: [1 2 4 5 10 20]
*/
