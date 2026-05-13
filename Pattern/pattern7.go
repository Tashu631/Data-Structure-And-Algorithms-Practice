package main

import "fmt"

func Pattern7() {
	for i := 1; i <= 5; i++ {
		for j := 5; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Print("*")
		}
		for j := 1; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")

	}
}

/* EXPLANATION - PATTERN 7:
Diamond Shape (Hollow Triangle/Pyramid)

Output:
    *
   **
  ***
 ****
*****

Logic:
- Outer loop (i): Runs from 1 to 5 for 5 rows
- For each row, 3 inner loops:

  1) First inner loop - Spaces (j = 5 down to i):
     - Prints (5-i) spaces to right-align the pattern
     - Row 1: prints 4 spaces
     - Row 2: prints 3 spaces
     - Row 3: prints 2 spaces, etc.

  2) Second inner loop - Leading stars (k = 0 to i):
     - Prints i stars (left side of diamond)

  3) Third inner loop - Trailing stars (j = 1 to i-1):
     - Prints (i-1) stars (right side of diamond)
     - Creates symmetry for diamond shape

Note: Total stars per row = i + (i-1) = 2i-1
*/
