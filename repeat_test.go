package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat('h', 10)
	want := "hhhhhhhhhh"

	assertCorrectMessage(t, got, want)
}

func ExampleRepeat() {
	res := Repeat('h', 5)
	fmt.Println(res)
	// Output: hhhhh
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat('a', 10)
	}
}
