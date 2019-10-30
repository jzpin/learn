package main

import (
	"context"
	"fmt"
)

// reproduce example from godoc
func main() {
	// example with cancel
	// create gen function
	gen := func(ctx context.Context) <-chan int {
		out := make(chan int)

		// start a go routine to send data
		go func() {
			n := 0
			for { // run forever until ctx cancel
				select {
				case <-ctx.Done():
					return
				case out <- n:
					n++
				}
			}
		}()

		return out
	}

	// create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// now we consume
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	// context with deadline

	// test other functions
	func1()
	func2()
	func3()
	func4timeout()
}
