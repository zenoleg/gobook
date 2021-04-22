package main

import "fmt"

func main() {
	data1 := []string{"a", "a", "b", "b", "b", "c", "c", "c", "c"}
	data2 := []string{"a", "b", "c"}
	data3 := []string{"a"}
	data4 := []string{}

	fmt.Println(removeNeighborDuplicates(data1))
	fmt.Println(removeNeighborDuplicates(data2))
	fmt.Println(removeNeighborDuplicates(data3))
	fmt.Println(removeNeighborDuplicates(data4))
}

func removeNeighborDuplicates(data []string) []string {
	if len(data) == 0 {
		return data
	}

	idx := 1

	for i := 1; i < len(data); i++ {
		if data[i] != data[i-1] {
			data[idx] = data[i]

			idx++
		}
	}

	return data[:idx]
}
