package main

import (
	"concurrency/chapter3/channels"
  "concurrency/chapter3/sync-impl/mutex"
  "concurrency/chapter3/sync-impl/waitgroups"
	"fmt"
	"runtime"
	"sync"
)

func main() {
  printHello()
  closureExample()
  printSliceWithoutParameter()
  printSliceWithParameter()
  memory()
  waitgroups.WaitGroupExample()
  waitgroups.Example2()
  mutex.IncDecr()

  channel.DeclareChannels()
  channel.HelloChannels()
  channel.DeadlockExample()  
  channel.ReadFromClosedChannel()

  channel.RangeOverChannel()
	channel.BufferedChannel()
  channel.ChannelSendRec()


}

func printHello() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("Hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
}

func closureExample() {
	var wg sync.WaitGroup
	salutation := "Hello"
	fmt.Println(&salutation)
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "Welcome"
		fmt.Println(&salutation)
	}()
	wg.Wait()
	fmt.Println(salutation)
}

func printSliceWithoutParameter() {
	var wg sync.WaitGroup
	salutations := []string{"hello", "good day", "greetings"}
	for _, salutation := range salutations {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

func printSliceWithParameter() {
	var wg sync.WaitGroup
	salutations := []string{"hello", "good day", "greetings"}
	for _, salutation := range salutations {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}

func memory() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c
	}

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	fmt.Println("before : ", before)
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}

	wg.Wait()
	after := memConsumed()
	fmt.Println("after : ", after)
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
	fmt.Println()
}
