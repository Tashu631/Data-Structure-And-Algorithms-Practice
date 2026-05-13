package main

import "fmt"

func main() {
	// Label each function output so we know whose answer it is
	fmt.Println("\n=== COUNT DIGITS ===")
	CountDigits()

	fmt.Println("\n=== REVERSE NUMBER ===")
	ReverseNumber()

	fmt.Println("\n=== PALINDROME CHECK ===")
	PalindromeNumber()

	fmt.Println("\n=== PRIME NUMBER CHECK ===")
	PrimeNumbers()

	fmt.Println("\n=== FACTORS ===")
	Factors()

	fmt.Println("\n=== GCD (GREATEST COMMON DIVISOR) ===")
	GCD()

	fmt.Println("\n=== ARMSTRONG NUMBER ===")
	ArmstrongNumber()
}

/* EXPLANATION:
This is the main entry point file for the math_questions package.
It calls all mathematical problem-solving functions in sequence:
- CountDigits(): Counts total digits in a number
- ReverseNumber(): Reverses a number
- PalindromeNumber(): Checks if a number reads same forwards and backwards
- PrimeNumbers(): Checks if a number is prime (only divisible by 1 and itself)
- Factors(): Finds all factors/divisors of a number
- GCD(): Finds Greatest Common Divisor of two numbers
- ArmstrongNumber(): Checks if sum of cubes of digits equals the number
*/
