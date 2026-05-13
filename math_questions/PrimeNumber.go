package main

import (
	"fmt"
	"math"
	"sort"
)

func PrimeNumbers() {
	num := 61
	count := 0
	var divisor []int

	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			divisor = append(divisor, i)
			count++
			if (num / i) != i {
				divisor = append(divisor, num/i)
				count++
			}
		}

	}
	sort.Ints(divisor)
	fmt.Println(divisor)
	if count == 2 {
		fmt.Println("number is a prime number")
	} else {
		fmt.Println("number is not  a prime number")
	}
}

/* EXPLANATION:
Prime Number Detection:

Prime Number: A natural number greater than 1 that has exactly 2 factors (1 and itself)
Examples: 2, 3, 5, 7, 11, 13, 17, 19, 23, 29...

Algorithm:
1. Find all divisors/factors of the number up to sqrt(num)
2. Count total divisors found
3. If count == 2, then it's prime (only factors are 1 and the number itself)
4. Otherwise, it's not prime

Why sqrt optimization?
- A prime number can only have divisors 1 and itself
- If any other divisor exists before sqrt(n), its complement exists after
- So checking only up to sqrt is sufficient

Example for 61:
- Divisors found: 1, 61 (count = 2)
- Output: "number is a prime number"

Example for 12:
- Divisors: 1, 2, 3, 4, 6, 12 (count = 6)
- Output: "number is not a prime number"

Time Complexity: O(√n)
*/

