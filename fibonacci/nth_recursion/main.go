package main

import "log"

/*
To find the \(n\)th Fibonacci number we can write code based on the mathematic formula
 for Fibonacci number \(n\):

\[F(n) = F(n-1) + F(n-2) \]
*/

func main() {
	log.Println(F(19))
}

func F(n int) int {
	if n <= 1 {
		return n
	} else {
		return F(n-1) + F(n-2)
	}
}
