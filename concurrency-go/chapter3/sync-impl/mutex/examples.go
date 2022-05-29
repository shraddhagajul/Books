package mutex

import (
	"fmt"
	"sync"
)

func IncDecr() {
	var count int
	var lock sync.Mutex

	increase := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Count : %v\n", count)
	}

	decrease := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Count : %v\n", count)
	}

	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increase()
		}()
	}

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrease()
		}()
	}

	wg.Wait()
}
