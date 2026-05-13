package main

import "fmt"

func Pattern1() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

}

/* EXPLANATION - PATTERN 1:
Square Pattern (5x5)

Output:
*****
*****
*****
*****
*****

Logic:
- Outer loop (i): Runs 5 times for 5 rows
- Inner loop (j): Runs 5 times for 5 columns
- Each iteration prints a star (*)
- After inner loop completes, newline is printed to move to next row

Time Complexity: O(n²) where n=5
Use Case: Basic practice for nested loops
*/
