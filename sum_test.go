package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3})
	want := 6

	assertCorrectMessage(t, got, want)
}

func sliceComp(t testing.TB, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GOT %v WANT %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{5, 6, 7})
	want := []int{6, 18}

	sliceComp(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("Slices with multiple items", func(t *testing.T) {
		got := SumAllTails([]int{2, 5}, []int{4, 5, 6})
		want := []int{5, 11}

		sliceComp(t, got, want)
	})

	t.Run("Slices with one item", func(t *testing.T) {
		got := SumAllTails([]int{2}, []int{4, 5, 6})
		want := []int{2, 11}

		sliceComp(t, got, want)
	})

	t.Run("Slices with no items", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 5, 6})
		want := []int{0, 11}

		sliceComp(t, got, want)
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3}, []int{5, 6, 7})
	}
}
