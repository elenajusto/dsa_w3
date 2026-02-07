package main

import "log"

/*
Recursion is when a function calls itself.

To implement the Fibonacci algorithm we need most of the same things as
in the code example above, but we need to replace the for loop with recursion.

To replace the for loop with recursion, we need to encapsulate much of the
code in a function, and we need the function to call itself to create a
new Fibonacci number as long as the produced number of Fibonacci numbers is
below, or equal to, 19.
*/

var count int

func main() {
	count = 2
	fibonacci(1, 0)
}

func fibonacci(prev1 int, prev2 int) {
	if count <= 19 {
		newFibo := prev1 + prev2
		log.Println("newFibo", newFibo)
		prev2 = prev1
		prev1 = newFibo
		count += 1
		fibonacci(prev1, prev2)
	} else {
		return
	}
}
