package main

import (
	"fmt"
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
}
