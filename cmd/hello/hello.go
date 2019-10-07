package main

import (
	"fmt"

	"github.com/jzpin/learn/pkg/mathutil"
	"github.com/jzpin/learn/pkg/stringutil"
)

func main() {
	fmt.Printf("Hello, world!\n")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	mathutil.FindPrime(100)
	//nums := []int{2, 2, 3, 4, 5}
	//mathutil.SquareNum(nums...)
}
