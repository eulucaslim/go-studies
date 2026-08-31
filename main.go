package main

import (
	"fmt"
	"go-studies/tempconv"
	"time"
)

func expensiveCall() {
	i, n := 0, 100000
	for i < n {
		i += 1
	}
	fmt.Printf("Finish! i value: %d", i)
}

func main() {
	start := time.Now()
	expensiveCall()
	elapsed := time.Since(start).Seconds()
	fmt.Printf("The call took %v to run.\n", elapsed)
	c := tempconv.Celsius(0)
	k := tempconv.Kelvin(373.15)
	fmt.Printf("Bruuuh! %v\n", tempconv.CToK(c))
	fmt.Printf("Kelvin to F: %v\n", tempconv.KToF(k))
}
