package main

import "fmt"

func main() {
	data := [4]int{1, 2, 3, 4}
	reverse(&data)

	fmt.Println(data)
}

func reverse(data *[4]int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
