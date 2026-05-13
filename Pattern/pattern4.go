package main

import "fmt"

func Pattern4() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Print("\n")
	}
}

/* EXPLANATION - PATTERN 4:
Right-Angled Triangle with Repeated Row Numbers

Output:
1
22
333
4444
55555

Logic:
- Outer loop (i): Runs from 1 to 5
- Inner loop (j): Runs from 1 to i
  - Row 1: prints 1 (once)
  - Row 2: prints 2 (twice)
  - Row 3: prints 3 (thrice)
  - Row 4: prints 4 (four times)
  - Row 5: prints 5 (five times)

Key Difference from Pattern3:
- Prints the row number (i), not the column number (j)
- Same number is repeated i times in each row
- Shows understanding of using outer loop variable in inner loop
*/
