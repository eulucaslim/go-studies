package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	result1 := printWithJoin(os.Args[1:]...)
	result2 := printWithoutJoin(os.Args[1:]...)

	fmt.Println("Result 1: " + result1)
	fmt.Println("Result 2: " + result2)
}

func printWithJoin(args ...string) string {
	start := time.Now()
	result := strings.Join(args, " ")
	fmt.Printf("Time print with Join: %.2f\n", time.Since(start).Seconds())
	return result
}

func printWithoutJoin(args ...string) string {
	start := time.Now()
	s, sep := "", " "

	for _, arg := range args {
		s += arg + sep
		sep = " "
	}
	fmt.Printf("Time print with Join: %.8f\n", time.Since(start).Seconds())
	return s
}
