package main

import "fmt"

func GCD() {
	a := 2
	b := 10

	for a >= 0 && b > 0 {
		if a > b {
			a = a % b
			// fmt.Print(a)
		} else {
			b = b % a
			// fmt.Print(b)
		}
	}
	if a == 0 {
		fmt.Println("GCD =", b)
	} else {
		fmt.Println("GCD =", a)
	}
}

