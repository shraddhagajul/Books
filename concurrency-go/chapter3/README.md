a goroutine is a function that is running concurrently.
Goroutines are unique to Go (though some other languages have a concurrency primitive that is similar). They’re not OS threads, and they’re not exactly green threads — threads that are managed by a language’s runtime — they’re a higher level of abstraction known as coroutines. Coroutines are simply concurrent subroutines (functions, closures, or methods in Go) that are nonpreemptive — that is, they cannot be interrupted. Instead, coroutines have multiple points throughout which allow for suspension or reentry.

Go’s mechanism for hosting goroutines is an implementation of what’s called an M:N scheduler, which means it maps M green threads to N OS threads. Goroutines are then scheduled onto the green threads. When we have more goroutines than green threads available, the scheduler handles the distribution of the goroutines across the available threads and ensures that when these goroutines become blocked, other goroutines can be run.

The go statement is how Go performs a fork, and the forked threads of execution are goroutines.
In order to a create a join point, you have to synchronize the main goroutine and the sayHello goroutine.

Closures close around the lexical scope they are created in, thereby capturing variables.


Chapter 2
Concurrency is a property of the code; parallelism is a property of the running program.
The first is that we do not write parallel code, only concurrent code that we hope will be run in parallel.
parallelism is a function of time, or context.

for example :if our context was a space of five seconds, and we ran two operations that each took a second to run, we would consider the operations to have run in parallel. If our context was one second, we would consider the operations to have run sequentially.


popular programming languages : If you had a lot of things you had to model concurrently and your machine couldn't handle that many threads, you created a thread pool and multiplexed your operations onto the thread pool.
it's common for languages to end their chain of abstraction at the level of the OS thread and memory access synchronization. Go takes a different route and supplants this with the concept of goroutines and channels.
Go has built a new set of concurrency primitives on top of the memory access synchronization primitives to provide you with an expanded set of things to work with.
