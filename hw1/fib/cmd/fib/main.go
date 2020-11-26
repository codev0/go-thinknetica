package main

import (
	"flag"
	"fmt"
	"go_thinknetica/hw1/fib/pkg/fib"
)

const MaxFibN = 20

func main() {
	var n = flag.Int("n", MaxFibN, "Number in Fibonacci sequence")

	flag.Parse()

	if *n > MaxFibN {
		fmt.Printf("Sorry, you cannot pass values more than: %v\n", MaxFibN)

		return
	}

	f := fib.Fib(*n)
	fmt.Printf("Fibbonacci number is: %v\n", f)
}
