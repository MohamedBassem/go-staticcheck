package pkg

import (
	"sync"
)

func fn() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()
	}()

	go func() {
		wg.Add(1) // MATCH /should call wg\.Add\(1\) before starting/
		wg.Done()
	}()

	wg.Add(1)
	go func(wg sync.WaitGroup) { // MATCH /should pass sync.WaitGroup by pointer/
		wg.Done()
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		wg.Done()
	}(&wg)

	wg.Wait()
}

func fn2(wg sync.WaitGroup) { // MATCH /should pass sync.WaitGroup by pointer/
	wg.Add(1)
}
