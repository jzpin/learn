package main

import (
	"fmt"
	"runtime"
	"sync"
)

// start two goroutines and print from a - z, A - Z

func main() {

	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup         // no need for new, just declare
	fmt.Println(runtime.NumCPU()) // 8 cores
	wg.Add(2)

	go func() {
		defer wg.Done()
		for c := 'a'; c <= 'z'; c++ {
			fmt.Printf("%c ", c)
		}
	}()

	go func() {
		defer wg.Done()
		for c := 'A'; c <= 'Z'; c++ {
			fmt.Printf("%c ", c)
		}
	}()

	wg.Wait()

}
