// package main

// import (
// 	"fmt"
// )

// type fibonacci struct {
// 	Iterator chan *int
// }

// func (f *fibonacci) Next() *int {
// 	return <-f.Iterator
// }

// func newFibonacci(limit int) *fibonacci {
// 	f := &fibonacci{
// 		make(chan *int, 1),
// 	}
// 	go func() {
// 		var current, next int
// 		for {
// 			if limit == 0 {
// 				close(f.Iterator)
// 				return
// 			}
// 			if current != 0 || next != 0 {
// 				next, current = current+next, next
// 			} else {
// 				next = 1
// 			}
// 			curr := current
// 			f.Iterator <- &curr
// 			limit--
// 		}
// 	}()
// 	return f
// }

// func main() {
// 	f := newFibonacci(10)
// 	for r := range f.Iterator {
// 		fmt.Println(*r)
// 	}
// }

package main

import "fmt"

func fibonacci(limit int) chan int {
	c := make(chan int)
	a := 0
	b := 1
	go func() {
		for {
			if limit == 0 {
				close(c)
				return
			}
			c <- a
			a, b = b, a+b
			limit--
		}
	}()
	return c
}

func main() {
	for r := range fibonacci(20) {
		fmt.Printf("%v ", r)
	}
}
