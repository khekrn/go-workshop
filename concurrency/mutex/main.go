package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	sc.value++
	sc.mu.Unlock()
}

func main() {
	sc := SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			sc.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Value:", sc.value)
}
