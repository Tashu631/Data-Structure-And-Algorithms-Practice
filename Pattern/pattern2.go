package main

import "fmt"

func Pattern2() {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

}

/* EXPLANATION - PATTERN 2:
Right-Angled Triangle (Increasing Stars)

Output:
*
**
***
****
*****

Logic:
- Outer loop (i): Runs 5 times for 5 rows (i = 0 to 4)
- Inner loop (j): Runs from 0 to i (inclusive)
  - Row 0: j runs 0 to 0 → 1 star
  - Row 1: j runs 0 to 1 → 2 stars
  - Row 2: j runs 0 to 2 → 3 stars
  - Row 3: j runs 0 to 3 → 4 stars
  - Row 4: j runs 0 to 4 → 5 stars

Key Point: Inner loop condition is j <= i (not j < i)
This makes each row have (i+1) stars
*/
