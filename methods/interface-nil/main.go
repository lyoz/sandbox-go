package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("M(): <nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	var t *T
	i = t

	fmt.Println("t==nil:", t == nil)
	fmt.Println("i==nil:", i == nil)

	// インターフェイスの指す具体的な値がnilの場合はnilチェックできる
	describe(i)
	i.M()
	i = &T{"Hello"}
	describe(i)
	i.M()

	// インターフェイス自体がnilだとランタイムエラー
	var j I
	describe(j)
	j.M()
}
