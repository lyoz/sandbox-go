package main

import (
	"fmt"
	"strings"
)

func testArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)
}

func testSlice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(primes)
	fmt.Println(s)

	s[0] = 1e9 + 7
	fmt.Println(primes)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(t)

	u := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(u[1:4])
	fmt.Println(u[:2])
	fmt.Println(u[1:])
	fmt.Println(u[:])
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func testLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	s = s[:0]
	printSlice(s)
	s = s[:4]
	printSlice(s)
	s = s[3:]
	printSlice(s)
}

func testNilSlice() {
	// 0要素のsliceはnilである
	var s []int
	fmt.Println(s, len(s), cap(s))
	fmt.Printf("%T\n", s)
	if s == nil {
		fmt.Println("nil!")
	}
}

func testMakeSlice() {
	a := make([]int, 5)
	printSlice(a)
	b := make([]int, 0, 5)
	printSlice(b)
	c := b[:2]
	printSlice(c)
	d := c[2:5]
	printSlice(d)
}

func testSliceOfSlice() {
	// tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func testAppendToSlice() {
	var s []int
	printSlice(s)
	s = append(s, 0)
	printSlice(s)
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func main() {
	// testArray()
	// testSlice()
	// testLenCap()
	// testNilSlice()
	// testMakeSlice()
	// testSliceOfSlice()
	testAppendToSlice()
}
