package main

import "fmt"

type fibonacci struct {
	current, next int
}

func (f *fibonacci) Next() int {
	if f.current != 0 || f.next != 0 {
		f.next, f.current = f.current+f.next, f.next
	} else {
		f.next = 1
	}
	return f.current
}

func main() {
	f := &fibonacci{}
	for i := 0; i < 10; i++ {
		fmt.Println(f.Next())
	}
}
