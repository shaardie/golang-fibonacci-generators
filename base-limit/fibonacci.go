package main

import "fmt"

type fibonacci struct {
	current, next, limit int
}

func (f *fibonacci) Next() *int {
	if f.limit == 0 {
		return nil
	}
	if f.current != 0 || f.next != 0 {
		f.next, f.current = f.current+f.next, f.next
	} else {
		f.next = 1
	}
	f.limit--
	return &f.current
}

func main() {
	f := &fibonacci{limit: 10}
	for {
		r := f.Next()
		if r == nil {
			return
		}
		fmt.Println(*r)
	}
}
