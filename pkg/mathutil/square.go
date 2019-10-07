package mathutil

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range nums {
			out <- i
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		for i := range in {
			result <- i * i
		}
		close(result)
	}()
	return result
}

// SquareNum return squares of number 1...n
func SquareNum(nums ...int) {
	seq := gen(nums...)
	out := sq(seq)
	for i := range out {
		fmt.Println(i)
	}
}
