package main

import (
	"errors"
	"net/http"
	"time"
)

func Racer(u1, u2 string) (string, error) {
	return CustomRacer(u1, u2, 10*time.Second)
}

func CustomRacer(u1, u2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(u1):
		return u1, nil
	case <-ping(u2):
		return u2, nil
	case <-time.After(timeout):
		return "", errors.New("response timeout")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
