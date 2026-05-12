package main

import (
	"fmt"
	"math"
)

func CountDigits() {
	num := 12345678910
	count := int(math.Log10(float64(num))) + 1
	fmt.Println("count digits:", count)
	// Process file content
}
