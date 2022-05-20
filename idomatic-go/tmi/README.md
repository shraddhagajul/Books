## Types 
We can only access the type from within its scope.
An abstract type is one that specifies what a type should do, but not how it is done. 
A concrete type specifies what and how. This means that it has a specified way to store its data and provides an implementation of any methods declared on the type.

## Methods
Method declarations look just like function declarations, with one addition: the receiver specification.
Just like functions, method names cannot be overloaded. You can use the same
method names for different types, but you can not use the same method name for two
different methods on the same type.

## Pointer Receiver and Value Receivers
The following rules help you determine when to use each
kind of receiver:
1. If your method modifies the receiver, you must use a pointer receiver.
2. If your method needs to handle nil instances, then it must use a pointer receiver.
3. If your method doesn't modify the receiver, you can use a value receiver.

Whether or not you use a value receiver for a method that doesn't modify the receiver
depends on the other methods declared on the type. When a type has any pointer
receiver methods, a common practice is to be consistent and use pointer receivers for
all methods, even the ones that don't modify the receiver.

runCounter()
we were able to call the pointer receiver method even though c is a value type. When you use a pointer receiver with a local variable that's a value type, Go automatically converts it to a pointer type. In this case, c.Increment() is converted to 
(&c).Increment().

updates()
the rules for passing values to functions still apply. If you pass
a value type to a function and call a pointer receiver method on the passed value, you
are invoking the method on a copy.

## Methods are functions too 
methodsSameAsFunc()

## Methods vs Func
The differentiator is whether or not your function depends on other data.
Any time our logic depends on values that are configured at startup or changed while your
program is running, those values should be stored in a struct and that logic should be
implemented as a method. If your logic only depends on the input parameters, then it
should be a function.

## type declarations aren't inheritance
var i int = 300
var s Score = 100
var hs HighScore = 200
hs = s // compilation error!
s = i // compilation error!
s = Score(i) // ok
hs = HighScore(s) // ok




