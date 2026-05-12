package main

import (
	"fmt"
	"math"
	"sort"
)

func Factors() {
	num := 20
	var divisor []int

	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			divisor = append(divisor, i)
		}
		if (num / i) != i {
			divisor = append(divisor, num/i)
		}
	}
	sort.Ints(divisor)
	fmt.Println(divisor)
}
