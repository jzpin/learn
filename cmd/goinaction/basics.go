package main

import (
	"fmt"
)

func main() {

	// 4.1 Arrays
	// - only initialize part of an array
	array1 := [5]int{1: 10, 2: 20}
	fmt.Println(array1)

	// - array of pointers, note pointer is *int
	array2 := [5]*int{1: new(int), 2: new(int)}
	// *array2[0] = 1 // this will panic
	*array2[1] = 2
	n := 3
	array2[3] = &n
	fmt.Println(array2)

	// - copy array
	var array3 [5]*int // cannot use := to declare empty array
	array3 = array2
	fmt.Println(array3) // copy array of pointer will refer to the same memory location

	// - use pointer
	var array4 [1000]int
	f1 := func(arr *[1000]int) { // nested function must be assigned to another variable
		fmt.Println("Inside f1")
	}
	f1(&array4)

	// 4.2 slice
	// slice can also be initialized with one element
	slice := []string{9: "a"}
	fmt.Println(slice)

	// different ways to make nil / empty slice
	// var slice []int
	// slice := []int {}
	zeroslice := make([]int, 0)
	fmt.Println(zeroslice)

	// experiment
	slice1 := make([]int, 5, 6)
	fmt.Println(cap(slice1))
	slice1 = append(slice1, 1) // reach max cap
	fmt.Println(cap(slice1))
	slice2 := slice1[2:5]
	fmt.Println(slice2)
	slice1[2] = 200
	fmt.Println(slice2)        // [2] modified
	slice1 = append(slice1, 6) // new slice is created
	slice1[2] = 100
	fmt.Println(slice2, cap(slice2)) // [2] not modified

	// use copy
	copied := make([]int, len(slice2), cap(slice2)*2) // double capacity
	copy(copied, slice2)
	slice2 = copied
	fmt.Println(copied, cap(slice2))

	// three index slice
	slice3 := slice2[0:1:2]
	fmt.Println(slice3, cap(slice3))

	// restrict cap
	s1 := []int{0, 1, 2, 3, 4}
	s2 := s1[0:3]
	fmt.Println(cap(s2)) // cap == 5!!
	s2 = append(s2, 100)
	fmt.Println(s1, s2)

	s3 := s1[0:3:3] // reduce capacity!!
	s3 = append(s3, 2000)
	fmt.Println(s1, s3)

	// 4.3 Map

}
