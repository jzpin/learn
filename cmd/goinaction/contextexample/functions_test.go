package main

import "testing"

func TestFunc1(t *testing.T) {
	func1()
	t.Error("show error")
}
