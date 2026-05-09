package main

import "fmt"

func Pattern1() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

}
