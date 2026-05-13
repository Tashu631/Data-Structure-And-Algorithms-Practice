package main

import "fmt"

func printNumbers(i, n int) {
	if i > n {
		return
	}
	fmt.Println(i)

	printNumbers(i+1, n)
}

/* EXPLANATION:
Linear Print 1 to N using Head Recursion

Goal: Print numbers from 1 to n using recursion (forward order)

Example: printNumbers(1, 4) outputs:
1
2
3
4

Algorithm:
1. Base Case: if i > n, return (stop recursion)
2. Print current number i
3. Recursively call with i+1

Execution Flow for printNumbers(1, 4):
Call Stack (Building phase):
  printNumbers(1,4) → prints 1, calls printNumbers(2,4)
                       printNumbers(2,4) → prints 2, calls printNumbers(3,4)
                                           printNumbers(3,4) → prints 3, calls printNumbers(4,4)
                                                               printNumbers(4,4) → prints 4, calls printNumbers(5,4)
                                                                                   printNumbers(5,4) → returns (BASE CASE)

Execution Timeline:
  1. printNumbers(1,4) prints 1, recurses
  2. printNumbers(2,4) prints 2, recurses
  3. printNumbers(3,4) prints 3, recurses
  4. printNumbers(4,4) prints 4, recurses
  5. printNumbers(5,4) hits base case, returns

Numbers print in FORWARD ORDER (1,2,3,4) because:
- We print BEFORE recursing (Head Recursion)
- Not waiting for deeper calls to complete

Key Characteristics:
- Parameters: i (current), n (limit)
- Base case: i > n
- Increment: i+1 with each call

Time Complexity: O(n)
Space Complexity: O(n) - call stack depth = n
*/
