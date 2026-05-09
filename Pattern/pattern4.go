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
