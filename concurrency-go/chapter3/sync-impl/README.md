printHello() => wg.Wait block the main goroutine until the goroutine hosting the sayHello function terminates. This is the join point.

closureExample() => that goroutines execute within the same address space they were created in, and so our program prints out the word "welcome".

printSliceWithoutParameter() => salutation becomes out of scope. Go runtime knows that reference to salutation variable is still being held. It is then transferred to the heap holding a reference to the last value in my string slice. 

The proper way to write this loop is to pass a copy of salutation into the closure so that by the time the goroutine is run, it will be operating on the data from its iteration of the loop:
printSliceWithParameter => we pass in the current iteration's variable to the closure. A copy of the string struct is made, thereby ensuring that when the goroutine is run, we refer to the proper string.

goroutines are not garbage collected with the runtime’s ability to introspect upon itself and measure the amount of memory allocated before and after goroutine creation:

Memory()

context switching, which is when something hosting a concurrent process must save its state to switch to running a different concurrent process. If we have too many concurrent processes, we can spend all of our CPU time context switching between them and never get any real work done. At the OS level, with threads, this can be quite costly.

we call Done using the defer keyword to ensure that before we exit the goroutine’s closure, we indicate to the WaitGroup that we’ve exited.


