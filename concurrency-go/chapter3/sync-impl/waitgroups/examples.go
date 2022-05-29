package waitgroups

import (
	"fmt"
	"sync"
)

func WaitGroupExample() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2st goroutine")
	}()

	wg.Wait()
	fmt.Println("All goroutines completed")
}

func Example2() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const greeters = 5

	var wg sync.WaitGroup
	wg.Add(greeters)
	for i := 0; i < greeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
