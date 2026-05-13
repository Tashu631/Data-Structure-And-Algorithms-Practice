// factorial of a number
package main

func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

/* EXPLANATION:
Factorial Calculation using Recursion

What is Factorial?
Factorial of n (written as n!) = n × (n-1) × (n-2) × ... × 2 × 1
Examples:
- 0! = 1 (by definition)
- 1! = 1
- 4! = 4 × 3 × 2 × 1 = 24
- 5! = 5 × 4 × 3 × 2 × 1 = 120

Recursive Formula:
factorial(n) = n × factorial(n-1)
factorial(0) = 1 (BASE CASE)
factorial(1) = 1 (BASE CASE)

Execution Example: factorial(4)
Call Stack (Building phase):
  factorial(4) → 4 * factorial(3)
                      factorial(3) → 3 * factorial(2)
                                      factorial(2) → 2 * factorial(1)
                                                      factorial(1) → returns 1
                                                      returns 2*1 = 2
                                      returns 3*2 = 6
                      returns 4*6 = 24

Result: factorial(4) = 24

Time Complexity: O(n)
Space Complexity: O(n) - due to call stack depth

Note: Factorial grows very quickly!
- 10! = 3,628,800
- 20! = 2,432,902,008,176,640,000
*/
