package main

import (
	"fmt"
	"math"
	"sort"
)

func PrimeNumbers() {
	num := 61
	count := 0
	var divisor []int

	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			divisor = append(divisor, i)
			count++
			if (num / i) != i {
				divisor = append(divisor, num/i)
				count++
			}
		}

	}
	sort.Ints(divisor)
	fmt.Println(divisor)
	if count == 2 {
		fmt.Println("number is a prime number")
	} else {
		fmt.Println("number is not  a prime number")
	}
}

