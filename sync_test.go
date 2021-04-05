package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Calling Inc() 3 times leaves the counter at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		want := 3

		assertCounter(t, &counter, want)
	})

	t.Run("It runs safely concurrently", func(t *testing.T) {
		counter := Counter{}
		wantedCount := 1000

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, &counter, wantedCount)
	})
}

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()

	if counter.Value() != want {
		t.Errorf("GOT %d WANT %d", counter.Value(), want)
	}
}
