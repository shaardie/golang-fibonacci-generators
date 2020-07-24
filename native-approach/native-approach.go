package main

import "fmt"

type fibonacci struct {
	a, b int
}

func (f *fibonacci) Next() int {
	r := f.a
	f.a, f.b = f.b, f.a+f.b
	return r
}

func main() {
	f := fibonacci{0, 1}
	for i := 0; i < 20; i++ {
		fmt.Printf("%v ", f.Next())
	}
}
