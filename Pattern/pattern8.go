package main

import "fmt"

func Pattern8() {
	for i := 5; i >= 1; i++ {
		for j := i;j>=0;j--{
			fmt.Print("*")
		}
		fmt.Print("\n")

	}
}
