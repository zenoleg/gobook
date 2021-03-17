// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		return
	}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		findFileWithDuplicates(f)
		_ = f.Close()
	}
}

func findFileWithDuplicates(f *os.File) {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++

		if counts[input.Text()] > 1 {
			fmt.Println("Found duplicates in", f.Name(), "file")

			return
		}
	}
}
