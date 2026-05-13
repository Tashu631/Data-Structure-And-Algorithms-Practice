package main

import "fmt"

func printString(n int) {
	if n == 5 {
		return
	}
	name := "tashu"
	fmt.Println(name)

	printString(n + 1)
}

/* EXPLANATION:
Print String N Times using Recursion

Goal: Print "tashu" multiple times without using a loop

Example: printString(0) outputs:
tashu
tashu
tashu
tashu
tashu
(5 times)

Algorithm:
1. Base Case: if n == 5, return (stop after 5 prints)
2. Print the string "tashu"
3. Recursively call with n+1 (increment counter)

Execution Flow for printString(0):
Call Stack (Building phase):
  printString(0) → prints "tashu", calls printString(1)
                   printString(1) → prints "tashu", calls printString(2)
                                    printString(2) → prints "tashu", calls printString(3)
                                                     printString(3) → prints "tashu", calls printString(4)
                                                                      printString(4) → prints "tashu", calls printString(5)
                                                                                       printString(5) → returns (BASE CASE)

Execution Timeline:
  1. n=0: prints "tashu"
  2. n=1: prints "tashu"
  3. n=2: prints "tashu"
  4. n=3: prints "tashu"
  5. n=4: prints "tashu"
  6. n=5: returns (BASE CASE)

Key Points:
- Counter 'n' starts at 0 and increments by 1 each call
- Base case checks if n == 5 (hardcoded limit)
- Prints BEFORE recursing (Head Recursion)

Comparison with Loop:
Loop version: for i := 0; i < 5; i++ { fmt.Println("tashu") }
Recursion: Achieves same result without explicit loop construct

Time Complexity: O(5) = O(1) - always prints 5 times
Space Complexity: O(5) = O(1) - call stack depth = 5
*/
