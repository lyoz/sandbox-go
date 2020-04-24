package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 200, 150))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
