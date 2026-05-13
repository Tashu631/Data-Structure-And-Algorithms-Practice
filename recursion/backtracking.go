// print lineraly from 1 to n , use n do without i
// backtracking
package main

import "fmt"

func print(n int) {
	if n < 1 {
		return
	}
	print(n - 1)
	fmt.Println(n)
}

/* EXPLANATION:
Backtracking Recursion - Print 1 to N

Goal: Print numbers 1 to n using recursion WITHOUT a loop variable

Example: print(4) outputs:
1
2
3
4

Algorithm:
1. Base Case: if n < 1, return (stop recursion)
2. Recursive Case:
   - First: Call print(n-1) to go deeper
   - Then: Print n after returning from recursion

Execution Flow for print(4):
Call Stack (building phase):
  print(4) → calls print(3)
             print(3) → calls print(2)
                        print(2) → calls print(1)
                                   print(1) → calls print(0)
                                              print(0) → returns (BASE CASE)

Execution Phase (backtracking):
  print(0) returns
  print(1) executes → prints 1
  print(2) executes → prints 2
  print(3) executes → prints 3
  print(4) executes → prints 4

Key Point:
- Recursion first goes to base case (n reaches 0)
- Then prints while returning back (backtracking)
- Numbers print in order 1,2,3,4 because of this backtracking

Time Complexity: O(n)
Space Complexity: O(n) - due to call stack
*/
