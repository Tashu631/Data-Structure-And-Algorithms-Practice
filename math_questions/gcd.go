package main

import "fmt"

func GCD() {
	a := 2
	b := 10

	for a >= 0 && b > 0 {
		if a > b {
			a = a % b
			// fmt.Print(a)
		} else {
			b = b % a
			// fmt.Print(b)
		}
	}
	if a == 0 {
		fmt.Println("GCD =", b)
	} else {
		fmt.Println("GCD =", a)
	}
}

/* EXPLANATION:
Greatest Common Divisor (GCD) using Euclidean Algorithm:

GCD: The largest number that divides both a and b
Example: GCD(10, 2) = 2, GCD(18, 24) = 6

Euclidean Algorithm - How it works:
1. Replace larger number with remainder of larger % smaller
2. Repeat until one number becomes 0
3. When one is 0, the other is the GCD

Step-by-step for a=2, b=10:
- Iteration 1: a < b, so b = 10 % 2 = 0
- Iteration 2: b = 0, so loop ends
- Result: GCD = a = 2

Why it works: GCD(a, b) = GCD(b, a mod b)
Time Complexity: O(log(min(a, b)))

Example: For a=2, b=10, output: "GCD = 2"
*/

