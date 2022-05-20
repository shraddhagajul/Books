## Pointers are variables that contain address of where another variable is stored

var x int32 = 10
pointerX = &x
Here (pointerX) is a variable. It contains address of where variable (x) is stored.

Zero value of pointer = nil (untyped identifier => represents lack of value of certain types)

(&) address operator => returns address of memory location where value is stored
(*) derefrencing operator => returns pointed-to value

x := 10
pointerX := &x 
fmt.Println(pointerX) //returns memory address
fmt.Prinln(*pointerX) //returns 10

## Before dereferencing a pointer make sure pointer is not nil, else it returns panic
var x *int // x is a pointer of with value 0
fmt.Println(x == nil) //returns true
fmt.Println(*x) //panic

var pointerToY *int 
y := 10
pointerToY = &y
Above 3 lines => pointerToY := &y 

## Cannot use (&) before a primitive literal. when we need a pointer to primitive type, declare a variable and point to it 

## Structs 
type Person struct {
	Firstname string
	Middlename *string
	Lastname string
}
Middlename (*string) => can be assigned values by following methods :
1. introduce a variable to hold content value 
2. PointerToPrimitiveType is a helper func to turn constant value into a pointer 

## Pointers and Immutability
PointersAndImmutability 
(f) is a pointer of type int. When failedUpdate is called, a copy of variable (f) is created. The copy (px) holds value of memory location of pointer (f).
x := 10
px = &x
Pointer (px) holds value of memory location where variable (x) is stored
This does not change value of pointer (f)

## Pointers and Mutuability
PointersAndMutability
(f) is a variable with value 2. We send a copy which contains address of variable (f) to func failedUpdate. It doesnt change the value of variable f.
However, in func update, (*px) => dereference => value at memory location px = 15
Hence value of f is changed

Ref : Page 116

## When returning values from a function, you should favor value types. Only use a
## pointer type as a return type if there is state within the data type that needs to be
## modified.

## Return pointer vs return value
If size of data structure is more ( > 10MB ) use pointers since pointers for all data structures have the same size.
If size of data structure is less, return value as it is better for performance.


## Zero vs No value assigned
If no values is assigned => pointer = nil
If pointer has zero value, we then can differentiate easily

if a nil pointer is passed into a function via a parameter or a field on a
parameter, you cannot set the value within the function as there is nowhere to store
the value. If a non-nil value is passed in for the pointer, do not modify it unless you
document the behavior.


## JSON

## Maps
any modifications made to a map that is passed to a function are reflected in the original variable that was passed in.This is because within the Go runtime, a map is implemented as a pointer to a struct. Passing a map to a function means that you are copying a pointer.

## Slices 
any modification to the contents of the slice is reflected in the original variable, but using append to change the length is not reflected in the original variable, even if the slice has a capacity greater than its length. That is because a slice is implemented as a struct

When a slice is copied to a different variable or passed to a function, a copy is made
of the length, capacity, and the pointer.

Changing the values in the slice changes the memory that the pointer points to, so the
changes are seen in both the copy and the original.

## Changes to the length and capacity are not reflected back in the original, because they are only in the copy. 
Changing the capacity means that the pointer is now pointing to
a new, bigger block of memory. Changing the [ capacity ] changes the storage

appendToSlice()
Changing the [ length ] is invisible in the original

## Reduce load on garbage collector
garbage => data that has no more pointers pointing to it => the memory that this data takes up can be reused

## Stack and Heap
A stack is a consecutive block of memory, and every function call in thread of
execution shares the same stack.
A stack pointer tracks the last location where memory was allocated; allocating additional memory is done by moving the stack pointer.

When a function is invoked, a new stack frame is created for the function's data. Local variables are stored on the stack, along with parameters passed into a function. Each new variable moves the stack pointer by the size of the value. When a function exits, its return values are copied back to the calling function via the stack and the stack pointer is moved back to the beginning of the stack frame for the exited function, deallocating all of the stack memory that was used by that function is local variables and parameters.

Go is unusual in that it can actually increase the size of a stack while the program is running. This is possible because each goroutine has its own stack and goroutines are managed by the Go runtime, not by the underlying operating system

To store something on the stack, you have to know exactly how big it is at compile
time.
If you know the size of memory to be allocated then its stored on stack else it is stored on heap

Ref : 124


## Java
In Java, local variables and parameters are stored in the stack, just like Go. However, as we discussed earlier, objects in Java are implemented as pointers. That means for every object variable instance, only the pointer to it is allocated on the stack; the data within the object is allocated on the heap. Only primitive values (numbers, booleans, and chars) are stored entirely on the stack.This means that the garbage collector in Java has to do a great deal of work. It also means that things like Lists in Java are actually a pointer to an array of pointers. Even though it looks like a linear data structure, reading it actually involves bouncing through memory, which is highly inefficient.

## Go encourages you to use pointers sparingly. 
We reduce the workload of the garbage collector by making sure that as much as possible is stored on the stack. Slices of structs or primitive types have their data lined up sequentially in memory for rapid access. And when the garbage collector does do work, it is optimized to return quickly rather than gather the most garbage. The key to making this approach work is to simply create less garbage in the first place. While focusing on
optimizing memory allocations can feel like premature optimization, the idiomatic
approach in Go is also the most efficient.

