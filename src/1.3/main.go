package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	benchmark("Slow", slow)
	benchmark("Fast", fast)
}

func benchmark(name string, function func()) {
	fmt.Println(name + " function benchmark:")

	start := time.Now()
	function()
	microseconds := time.Since(start).Microseconds()

	fmt.Println(name+" function execution time:", microseconds, "microseconds")
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
