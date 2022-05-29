a rendezvous point for goroutines waiting for or announcing the occurrence
of an event.
The NewCond function takes in a type that satisfies the sync.Locker interface. This is what allows the Cond type to facilitate coordination with other goroutines in a concurrent-safe way.
to Wait doesn’t just block, it suspends the current goroutine, allowing other goroutines to run on the OS thread.

upon entering Wait, Unlock is called on the Cond variable’s Locker, and upon exiting Wait, Lock is called on the Cond variable’s Locker.

