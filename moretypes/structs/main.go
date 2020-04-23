package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	vp = &Vertex{10, 20}
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	p := &v
	(*p).X = 100
	p.Y = 200
	fmt.Println(v)

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(vp)
}
