package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.value++
	time.Sleep(100 * time.Millisecond)
}

func main() {
	sc := SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			sc.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Value:", sc.value)
}
