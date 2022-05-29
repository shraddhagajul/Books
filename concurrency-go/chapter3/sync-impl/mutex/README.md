A Mutex provides a concurrent-safe way to express exclusive access to these shared resources. To borrow a Goism, whereas channels share memory by communicating, a Mutex shares memory by creating a convention developers must follow to synchronize access to the memory.
You are responsible for coordinating access to this memory by guarding access to it with a mutex.

sync.RWMutex is conceptually the same thing as a Mutex: it guards access to memory; however, RWMutex gives you a little bit more control over the memory. You can request a lock for reading, in which case you will be granted access unless the lock is being held for writing. This means that an arbitrary number of readers can hold a reader lock so long as nothing else is holding a writer lock.
