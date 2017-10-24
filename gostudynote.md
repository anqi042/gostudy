##Go's declaration syntax  
```
x: int
p: pointer to int
a: array[3] of int
```
Go's declarations read left to right.In the interests of brevity it drops the colon and removes some of the keywords  

```   
x int
p *int
a [3]int   
```

> https://blog.golang.org/gos-declaration-syntax

##SLICE  
Slices are like references to arrays
The length of a slice is the number of elements it contains.
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
A nil slice has a (1)length and (2)capacity of 0 and (3)has no underlying array. ---slice只要指向一个array，即使将一个slice元素都drop掉后，也不满足条件3，不能称为nil slice


##MAP
The zero value of a map is?nil. A?nil?map has no keys, nor can keys be added.
The **make** function returns a map of the given type, initialized and ready for use.
---比如定义一个var m map[string]string  ，m就是nil的，需要make后才能使用

go中的map是hashmap，是一种无序map，linkedmap或者treemap可以实现有序map。

Maps are not safe for **concurrent** use
https://blog.golang.org/go-maps-in-action
---需要加mutex
---golang1.9 sync包中增加新的map类型，支持并发，多个goroutine可以安全并发的调用它。


##RANGE
When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
---由于range value是copy，使用range  value指针赋值是错误的

##METHODS
Remember: a method is just a function with a receiver argument.
You can declare a method on non-struct types, too.
You can only declare a method with a receiver whose type is defined in the same package as the method(includes the built-in types such as?int)
---所以要将内建类型进行重定义：type MyFloat float64

There are two reasons to use a pointer receiver.
the method can modify the value that its receiver points to.
avoid copying the value on each method call

All methods on a given type should have either value or pointer receivers, but not a mixture of both


##INTERFACES
An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.

Interface values can be thought of as a tuple of a value and a concrete type:
(value, type)

An interface value that holds a nil concrete value is itself non-nil.
---example: https://tour.golang.org/methods/12

A nil interface value holds neither value nor concrete type.

An empty interface may hold values of any type.



##Stringers 没看懂???
One of the most ubiquitous interfaces is?Stringer?defined by the?fmt?package.
type Stringer interface {
 
    String() string
}
}
A Stringer is a type that can describe itself as a string. The?fmt?package (and many others) look for this interface to print values.
https://tour.golang.org/methods/17



##GOROUTINES
It's an independently executing function, launched by a go statement.
It has its own call stack, which grows and shrinks as required.
It's not a thread.
Go supports concurrency---In the language and runtime, not a library.This changes how you structure your programs.
Concurrency is not parallelism---golang.org/s/concurrency-is-not-parallelism
Rough analogy: writing to a file by name (process, Erlang) vs. writing to a file descriptor (channel, Go).


##CHANNEL
Note:?Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a?range?loop.
Don't communicate by sharing memory, share memory by communicating.

而channel，可以理解为：用于并发单元间的数据解耦的、阻塞的、带类型的、并发安全的消息队列。

##SELECT
- All channels are evaluated.?
- Selection blocks until one communication can proceed, which then does.?
- If multiple can proceed, select chooses pseudo-randomly.?
- A default clause, if present, executes immediately if no channel is ready.


##Concurrency Example:
google search优化过程：  
(1)串行搜索（1.0）  
(2)并行搜索，主进程收齐结果返回（2.0）  
(3)设定超时时间，超时放弃等待（2.1）  
(4)使用多个服务器副本避免超时（3.0）    
Summary
In just a few simple transformations we used Go's concurrency primitives to convert a slow sequential failure-sensitive program into one that is fast concurrent replicated robust.

More party tricks  
There are endless ways to use these tools, many presented elsewhere.  
Chatroulette toy:  
>golang.org/s/chat-roulette  

Load balancer:   
>golang.org/s/load-balancer   

Concurrent prime sieve:   

>golang.org/s/prime-sieve  

Concurrent power series (by McIlroy):

>golang.org/s/power-series

Conclusions
Goroutines and channels make it easy to express complex operations dealing with  
multiple inputs  
multiple outputs  
timeouts  
failure  
And they're fun to use.


##MAKE vs NEW（https://golang.org/doc/effective_go.html#allocation_new）  
"new" does not?initialize?the memory, it only?zeros?it.?  
new(T)?allocates zeroed storage for a new item of type?T?and returns its address  
"make" creates slices, maps and channels only,and it returns an?initialized?(not?  zeroed) value of type?T?(not?*T).  

##return local variable address in go(effective go)
it's perfectly OK to return the address of a local variable; the storage associated with the variable survives after the function returns. In fact, taking the address of a composite literal allocates a fresh instance each time it is evaluated
