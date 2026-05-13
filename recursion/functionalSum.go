// sum of first n numbers
// functional
package main

func sumFunc(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumFunc(n-1)
}

/* EXPLANATION:
Sum of First N Numbers - Functional Recursion

Goal: Calculate sum 1 + 2 + 3 + ... + n
Example: sumFunc(5) = 1 + 2 + 3 + 4 + 5 = 15

Recursive Formula:
sumFunc(n) = n + sumFunc(n-1)
Base Case: sumFunc(0) = 0

Execution Example: sumFunc(5)
Call Chain:
  sumFunc(5) = 5 + sumFunc(4)
  sumFunc(4) = 4 + sumFunc(3)
  sumFunc(3) = 3 + sumFunc(2)
  sumFunc(2) = 2 + sumFunc(1)
  sumFunc(1) = 1 + sumFunc(0)
  sumFunc(0) = 0 (BASE CASE)

Backtracking (returning values):
  sumFunc(0) = 0
  sumFunc(1) = 1 + 0 = 1
  sumFunc(2) = 2 + 1 = 3
  sumFunc(3) = 3 + 3 = 6
  sumFunc(4) = 4 + 6 = 10
  sumFunc(5) = 5 + 10 = 15

This is called "Functional Recursion" because:
- Each call returns a value that's used in the previous call
- It builds the result from multiple return values

Mathematical Formula: Sum = n*(n+1)/2
For n=5: 5*6/2 = 15 ✓

Time Complexity: O(n)
Space Complexity: O(n) - call stack

Advantage over loop: More elegant and expresses mathematical thinking
*/
