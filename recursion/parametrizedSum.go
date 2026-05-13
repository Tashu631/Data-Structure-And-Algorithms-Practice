// sum of first n numbers
// parametrised
package main

import "fmt"

func sumnumbers(i, sum int) {
	if i < 1 {
		fmt.Println(sum)
		return
	}
	sumnumbers(i-1, sum+i)
}

/* EXPLANATION:
Sum of First N Numbers - Parametrized Recursion

Goal: Calculate sum 1 + 2 + 3 + ... + n (n=4 → sum=10)

Example: sumnumbers(4, 0) → output: 10

Recursive Formula:
sumnumbers(i, sum) = sumnumbers(i-1, sum+i)
Base Case: if i < 1, print sum and return

Key Difference from Functional Recursion:
- Parametrized: Accumulates result in a parameter (sum)
- Functional: Calculates result from return values

Execution Example: sumnumbers(4, 0)
Call Chain (Forward execution):
  sumnumbers(4, 0) → calls sumnumbers(3, 4)  [sum = 0+4 = 4]
                     sumnumbers(3, 4) → calls sumnumbers(2, 7)  [sum = 4+3 = 7]
                                        sumnumbers(2, 7) → calls sumnumbers(1, 9)  [sum = 7+2 = 9]
                                                           sumnumbers(1, 9) → calls sumnumbers(0, 10)  [sum = 9+1 = 10]
                                                                               sumnumbers(0, 10) → BASE CASE
                                                                                                      Prints 10 and returns

Parameter Progression:
  i=4, sum=0
  i=3, sum=4   (0 + 4)
  i=2, sum=7   (4 + 3)
  i=1, sum=9   (7 + 2)
  i=0, sum=10  (9 + 1)
  i<0, print 10 (BASE CASE)

Why "Parametrized"?
- Accumulation happens through function parameters
- No need to combine return values
- Result is built incrementally as we go deeper

Advantage: This approach is "Tail Recursive"
- Can be optimized by compiler (tail call optimization)
- Doesn't build up return values on stack

Time Complexity: O(n)
Space Complexity: O(n) - call stack (but can be optimized to O(1))

Mathematical Formula: Sum = n*(n+1)/2
For n=4: 4*5/2 = 10 ✓
*/
