package main

import "fmt"

func Pattern5() {
	for i := 5; i >0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

/* EXPLANATION - PATTERN 5:
Inverted Triangle (Decreasing Stars) - Reverse of Pattern2

Output:
*****
****
***
**
*

Logic:
- Outer loop (i): Runs from 5 DOWN TO 1 (i-- instead of i++)
  - Row 1: i=5, j runs 1 to 5 → 5 stars
  - Row 2: i=4, j runs 1 to 4 → 4 stars
  - Row 3: i=3, j runs 1 to 3 → 3 stars
  - Row 4: i=2, j runs 1 to 2 → 2 stars
  - Row 5: i=1, j runs 1 to 1 → 1 star
- Inner loop (j): Runs from 1 to i

Key Difference:
- Loop condition: i > 0 (not i < 5)
- Loop decrement: i-- (not i++)
- Creates pyramid that narrows downward
*/
