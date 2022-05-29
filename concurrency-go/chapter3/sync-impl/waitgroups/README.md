Add with an argument of 1 to indicate that one goroutine is beginning.
we call Done using the defer keyword to ensure that before we exit the goroutine's closure, we indicate to the WaitGroup that we've exited.
we call Wait, which will block the main goroutine until all goroutines have indicated they have exited.
calls to Add increment the counter by the integer passed in, and calls to Done decrement the counter by one. Calls to Wait block until the counter is zero.

Notice that the calls to Add are done outside the goroutines they're helping to track. If we didnt do this, we would have introduced a race condition, because remember from "Goroutines" that we have no guarantees about when the goroutines will be scheduled; we could reach the call to Wait before either of the goroutines begin. Had the calls to Add been placed inside the goroutines closures, the call to Wait could have returned without blocking at all because the calls to Add would not have taken place. Its customary to couple calls to Add as closely as possible to the goroutines theyre helping to track,

