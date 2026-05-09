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
