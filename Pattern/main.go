package main

import "fmt"

func main() {
	// Label each pattern so we know which one is printing
	fmt.Println("\n===== PATTERN 1: SQUARE (5x5) =====")
	Pattern1()

	fmt.Println("\n===== PATTERN 2: RIGHT-ANGLED TRIANGLE (INCREASING) =====")
	Pattern2()

	fmt.Println("\n===== PATTERN 3: TRIANGLE WITH NUMBERS =====")
	Pattern3()

	fmt.Println("\n===== PATTERN 4: TRIANGLE WITH REPEATED ROW NUMBER =====")
	Pattern4()

	fmt.Println("\n===== PATTERN 5: INVERTED TRIANGLE =====")
	Pattern5()

	fmt.Println("\n===== PATTERN 6: INVERTED TRIANGLE WITH NUMBERS =====")
	Pattern6()

	fmt.Println("\n===== PATTERN 7: DIAMOND SHAPE =====")
	Pattern7()

}

/* EXPLANATION:
Pattern Printing Programs - Main Entry Point

This file contains the main() function that calls 8 different pattern functions.
Each pattern demonstrates nested loops with different logic:

Pattern1: Square of stars (5x5)
Pattern2: Right-angled triangle (increasing stars)
Pattern3: Right-angled triangle with numbers 1,2,3...
Pattern4: Right-angled triangle with repeated row number
Pattern5: Inverted triangle of stars
Pattern6: Inverted triangle with numbers
Pattern7: Diamond shape (hollow)


These programs help practice:
- Nested loops (outer loop for rows, inner loop for columns)
- Loop control with i++ and i--
- String/character printing
- Understanding loop iteration patterns
*/
