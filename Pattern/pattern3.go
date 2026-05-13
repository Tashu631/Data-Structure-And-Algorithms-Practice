package main

import "fmt"

func Pattern3() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Print("\n")
	}

}

/* EXPLANATION - PATTERN 3:
Right-Angled Triangle with Numbers (1,2,3...)

Output:
1
12
123
1234
12345

Logic:
- Outer loop (i): Runs from 1 to 5 (i starts at 1, not 0)
- Inner loop (j): Runs from 1 to i
  - Row 1: j prints 1
  - Row 2: j prints 1, 2
  - Row 3: j prints 1, 2, 3
  - Row 4: j prints 1, 2, 3, 4
  - Row 5: j prints 1, 2, 3, 4, 5

Key Difference from Pattern2:
- Both outer and inner loops start from 1 (not 0)
- Inner loop prints the loop variable j, not a fixed character
- Creates increasing number sequences
*/
