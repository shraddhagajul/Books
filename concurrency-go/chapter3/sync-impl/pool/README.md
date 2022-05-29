Pool is a concurrent-safe implementation of the object pool pattern.
At a high level, a the pool pattern is a way to create and make available a fixed number, or pool, of things for use. It’s commonly used to constrain the creation of things that are expensive (e.g., database connections) so that only a fixed number of them are ever created, but an indeterminate number of operations can still request access to these things. In the case of Go’s sync.Pool, this data type can be safely used by multiple goroutines.

memory needed for the resource is reduced ref:Loc:1743
Another common situation where a Pool is useful is for warming a cache of pre-allocated objects for operations that must run as quickly as possible. In this case, instead of trying to guard the host machine’s memory by constraining the number of objects created, we’re trying to guard consumers’ time by front-loading the time it takes to get a reference to another object. This is very common when writing high-throughput network servers that attempt to respond to requests as quickly as possible.

