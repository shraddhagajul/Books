package channel

import (
	"fmt"
	"sync"
)

func DeclareChannels() {
	var datastream chan interface{}
	fmt.Println(datastream) //returns nil
	datastream = make(chan interface{})
	fmt.Println(datastream) // returns in memory address

read := make(<-chan interface{})
fmt.Println(read)

send := make(chan<- interface{})
fmt.Println(send)

intChannel := make(chan<- int)
fmt.Println(intChannel)
}

func HelloChannels() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello Channels!"
	}()
	fmt.Println(<-stringStream)
}

// The DeadlockExample goroutine is waiting for a value to be placed onto the stringStream channel, 
// and because of our conditional, this will never happen. When the anonymous goroutine exits, 
// Go correctly detects that all goroutines are asleep, and reports a deadlock.

func DeadlockExample() {
	stringStream := make(chan string)
	go func() {
		if 0 != 1 {
			return
		}
		stringStream <- "Hello channels"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("%v : %v",salutation, ok)
}

func ReadFromClosedChannel() {
	intStream := make(chan int)
	close(intStream)
	intVal, ok := <-intStream
	fmt.Printf("%v : %v\n",intVal, ok)
}

func RangeOverChannel() {
	intStream := make(chan int)
	go func() {
	defer close(intStream)
		for i:=0; i<5; i++ {
			intStream <- i+2
		}
	}()


	for val := range intStream {
		fmt.Println("integer : ", val)
	}
}

func BufferedChannel() {
	runeStream := make(chan rune, 4)
	fmt.Println(cap(runeStream))
	
	var wg sync.WaitGroup;
	wg.Add(1)

	go func() {
		defer close(runeStream)
		defer wg.Done()
		runeStream<- 'A'
		runeStream<- 'B'
		runeStream<- 'C'
		runeStream<- 'D'
	}()

	wg.Wait()
	fmt.Println(cap(runeStream))
}

func ChannelSendRec() {
	chanOwner := func() <-chan int {
		receiverChan := make(chan int, 2) //1
		go func() { //2
			defer close(receiverChan) //3
			for i:=0;i<5;i++ { 
				receiverChan<-i+1 
			}
		}()
		return receiverChan
	}

	resultStream := chanOwner()
	for value := range resultStream {
		fmt.Printf("Receieved : %d\n", value)
	}
	fmt.Println("Done receieving")
}