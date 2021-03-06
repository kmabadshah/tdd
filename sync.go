package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.count++
	defer c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.count
}
