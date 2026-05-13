package main

import "fmt"

func main() {
	// Label each recursion function output clearly
	fmt.Println("\n===== FIBONACCI(6) =====")
	fmt.Println("Answer:", fibo(6))
	
	fmt.Println("\n===== PRINT STRING 5 TIMES (RECURSION) =====")
	n := 0
	printString(n)
	
	fmt.Println("\n===== PRINT NUMBERS 1 TO 4 (HEAD RECURSION) =====")
	n = 4
	printNumbers(1, n)
	
	fmt.Println("\n===== PRINT 4 DOWN TO 1 (BACKTRACKING) =====")
	printNto1(n, n)
	
	fmt.Println("\n===== PRINT 1 TO 4 LINEARLY (BACKTRACKING) =====")
	print(n)
	
	fmt.Println("\n===== SUM OF 1 TO 4 (PARAMETRIZED RECURSION) =====")
	sum := 0
	sumnumbers(n, sum)
	
	fmt.Println("\n===== SUM OF 1 TO 4 (FUNCTIONAL RECURSION) =====")
	fmt.Println("Answer:", sumFunc(n))
	
	fmt.Println("\n===== FACTORIAL(4) =====")
	fmt.Println("Answer:", factorial(n))

	fmt.Println("\n===== REVERSE ARRAY =====")
	arr := []int{90, 2877, 67, 92}
	fmt.Println("Original:", []int{90, 2877, 67, 92})
	l := 0
	r := len(arr) - 1
	fmt.Println("Reversed:", reverseArray(arr, l, r))

	fmt.Println("\n===== PALINDROME CHECK =====")
	fmt.Println("Is 'racecar' palindrome?", isPalindrome("racecar", 0, 6))
	fmt.Println("Is 'hello' palindrome?", isPalindrome("hello", 0, 4))

}
