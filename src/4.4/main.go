package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(data, 3)

	fmt.Println(data)
}

func reverse(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func rotate(data []int, n int) {
	reverse(data[:n])
	reverse(data[n:])
	reverse(data)
}
