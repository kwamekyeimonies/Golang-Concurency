When a goroutine sends/receives a value on a channel, it blocks until another go routine sends or receives a value on the same
Channel.
Atomicity
Race conditions->When multiple concurrent programs try to access or share data
Deadlocks
Livelocks
Starvation
Code Complexity

Always Pass a Waitgroup as a pointer

What is a Goroutine?
Is a lightweight execution thread in the Go programming Language and function that executes concurrently with the rest of the program. Dealing with things at the same time.

Thread -->Is a program designed to be scheduled and exeuted by the CPU 


FORK-JOIN MODEL

This shows that,the efficiency(speed) of a programs is not solely determined by How Efficient the Algorithm is,but also the Operations been performed by the CPU at the same time.

The WaitGroup can be used interchangeably with the Channels as Demonstrated in this Project

Problems Faced with Concurrency

-Race Conditions
Multiples pieces of concurrent program shares data
-Deadlocks
where processes wait for one another
-Livelocks
-Starvation
-Code Complexity