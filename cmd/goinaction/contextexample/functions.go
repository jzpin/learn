package main

import (
	"fmt"
	"time"
)

// boring example from go concurrency patterns
// boring function send char from msg to outside
func boring1(msg string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%v say: %d", msg, i)
	}
}

// print 5 msgs
func func1() {
	ch := make(chan string)
	go boring1("boring", ch)
	for i := 0; i < 5; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
}

// pattern 1 return a chan
func boring2(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%v say again: %d", msg, i)
		}
	}()
	return ch
}

func func2() {
	ch := boring2("boring")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

// fan in function, merge two channel into one
func fanin(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func func3() {
	input1 := boring2("J")
	input2 := boring2("K")
	ch := fanin(input1, input2)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func func4timeout() {
	endless := boring2("kite")
	timeout := time.After(1 * time.Second)
outloop: // name a loop
	for {
		select {
		case <-timeout:
			break outloop // this is interesting! like break version goto
		case s := <-endless:
			fmt.Println(s)
		}
	}
	fmt.Println("TIMEOUT")
}
