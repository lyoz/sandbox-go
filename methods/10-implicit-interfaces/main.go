package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// "T implements I" のような構文はない
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
