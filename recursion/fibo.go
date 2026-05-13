package main

func fibo(n int) int {

	if n <= 1 {

		return n

	}

	last := fibo(n - 1)

	secondLast := fibo(n - 2)

	return last + secondLast

}

/* EXPLANATION:
Fibonacci Sequence using Recursion

What is Fibonacci Sequence?
Each number is the sum of the two preceding numbers:
0, 1, 1, 2, 3, 5, 8, 13, 21, 34...

Recursive Formula:
fibo(n) = fibo(n-1) + fibo(n-2)
Base Case: fibo(0) = 0, fibo(1) = 1

Examples:
- fibo(0) = 0
- fibo(1) = 1
- fibo(2) = fibo(1) + fibo(0) = 1 + 0 = 1
- fibo(3) = fibo(2) + fibo(1) = 1 + 1 = 2
- fibo(4) = fibo(3) + fibo(2) = 2 + 1 = 3
- fibo(5) = fibo(4) + fibo(3) = 3 + 2 = 5
- fibo(6) = fibo(5) + fibo(4) = 5 + 3 = 8

Execution Tree for fibo(5):
         fibo(5)
        /      \
    fibo(4)  fibo(3)
    /  \      /  \
  fibo(3) fibo(2) fibo(2) fibo(1)
  ...and so on...

PROBLEM: Exponential Time Complexity O(2^n)
- fibo(5) calls fibo(3) multiple times
- fibo(3) calls fibo(1) and fibo(2) multiple times
- Lot of redundant calculations!

Time Complexity: O(2^n) - VERY SLOW for large n
Space Complexity: O(n) - call stack depth

Better Approaches:
1. Memoization: Cache results to avoid recalculation
2. Dynamic Programming: Build solution bottom-up
3. Iterative: Use loop instead of recursion

For n=40+, recursion is TOO SLOW!
*/
