package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	slowBenchmark()
	fastBenchmark()
}

func slowBenchmark() {
	fmt.Println("Slow function benchmark:")

	startSlow := time.Now()
	slow()
	slowTime := time.Since(startSlow).Microseconds()

	fmt.Println("Slow function execution time:", slowTime, "ms")
}

func fastBenchmark() {
	fmt.Println("Fast function benchmark:")

	startSlow := time.Now()
	fast()
	fastTime := time.Since(startSlow).Microseconds()

	fmt.Println("Fast function execution time:", fastTime, "ms")
}

func slow() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}

func fast() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
