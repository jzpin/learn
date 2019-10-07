package mathutil

import "fmt"

func gen2ToN(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 2; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func isPrime(done chan struct{}, in <-chan int, result chan<- int) {
	s, ok := <-in
	if ok == false { // closed
		close(done)
		return
	}
	result <- s
	out := make(chan int)
	go func() {
		for i := range in {
			if i%s != 0 {
				out <- i
			}
		}
		close(out)
	}()
	done2 := make(chan struct{})
	go isPrime(done2, out, result)
	<-done2
	close(done)
}

// FindPrime will print prime number smaller than n
func FindPrime(n int) {
	nums := gen2ToN(n)
	result := make(chan int)
	done := make(chan struct{})
	go isPrime(done, nums, result)
	for {
		select {
		case v := <-result:
			fmt.Println(v)
		case <-done:
			return
		}
	}
}
