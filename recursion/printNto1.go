package main

import "fmt"

func printNto1(i, n int) {
	if i < 1 {
		return
	}
	fmt.Println(i)
	printNto1(i-1, n)
}

/* EXPLANATION:
Print N Down To 1 using Head Recursion

Goal: Print numbers from n down to 1 in reverse order

Example: printNto1(5, 5) outputs:
5
4
3
2
1

Algorithm:
1. Base Case: if i < 1, return (stop recursion)
2. Print current number i
3. Recursively call with i-1 (decrement counter)

Execution Flow for printNto1(5, 5):
Call Stack (Building phase):
  printNto1(5,5) → prints 5, calls printNto1(4,5)
                   printNto1(4,5) → prints 4, calls printNto1(3,5)
                                    printNto1(3,5) → prints 3, calls printNto1(2,5)
                                                     printNto1(2,5) → prints 2, calls printNto1(1,5)
                                                                      printNto1(1,5) → prints 1, calls printNto1(0,5)
                                                                                       printNto1(0,5) → returns (BASE CASE)

Execution Timeline:
  1. i=5: prints 5, recurses with i=4
  2. i=4: prints 4, recurses with i=3
  3. i=3: prints 3, recurses with i=2
  4. i=2: prints 2, recurses with i=1
  5. i=1: prints 1, recurses with i=0
  6. i=0: hits base case (i < 1), returns

Numbers print in DESCENDING ORDER (5,4,3,2,1) because:
- We print BEFORE recursing (Head Recursion)
- Decrement i with each recursive call

Comparison with printNumbers:
- printNumbers(1, n): prints 1 to n (ascending) by incrementing
- printNto1(n, n): prints n to 1 (descending) by decrementing

Key Insight:
- First parameter 'i' is the counter
- Second parameter 'n' is limit (not really used, just passed along)
- Base case: i < 1 (stops when counter reaches 0)

Time Complexity: O(n)
Space Complexity: O(n) - call stack depth = n
*/
