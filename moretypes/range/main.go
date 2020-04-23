package main

import "fmt"

const N = 10

func main() {
	pow := make([]int, N)
	for i := range pow {
		pow[i] = 1 << i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
