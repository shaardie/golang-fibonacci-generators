package main

import "fmt"

type fibonacci struct {
	a, b, limit int
}

func (f *fibonacci) Next() *int {
	if f.limit == 0 {
		return nil
	}
	r := f.a
	f.a, f.b = f.b, f.a+f.b
	f.limit--
	return &r
}

func main() {
	f := fibonacci{0, 1, 20}
	for {
		r := f.Next()
		if r == nil {
			return
		}
		fmt.Printf("%v ", *r)
	}
}
