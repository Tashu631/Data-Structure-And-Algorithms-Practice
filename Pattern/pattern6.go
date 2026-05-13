package main

import "fmt"

func Pattern6() {
	for i := 5; i >0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Print("\n")
	}
}

/* EXPLANATION - PATTERN 6:
Inverted Triangle with Numbers (Reverse of Pattern3)

Output:
12345
1234
123
12
1

Logic:
- Outer loop (i): Runs from 5 DOWN TO 1 (i--)
  - Row 1: i=5, prints 1,2,3,4,5
  - Row 2: i=4, prints 1,2,3,4
  - Row 3: i=3, prints 1,2,3
  - Row 4: i=2, prints 1,2
  - Row 5: i=1, prints 1
- Inner loop (j): Runs from 1 to i, printing j each time

Key Features:
- Outer loop decrements (i--) instead of increment
- Inner loop always prints incrementing numbers (1 to i)
- Combines inverted pattern with numeric output
*/
