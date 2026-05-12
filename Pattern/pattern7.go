package main

import "fmt"

func Pattern7() {
	for i := 1; i <= 5; i++ {
		for j := 5; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Print("*")
		}
		for j := 1; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")

	}
}
