package main

import "log"

/*
   Two variables to hold the previous two Fibonacci numbers
   A for loop that runs 18 times
   Create new Fibonacci numbers by adding the two previous ones
   Print the new Fibonacci number
   Update the variables that hold the previous two fibonacci numbers
*/

func main() {
	prev2 := 0
	prev1 := 1

	log.Println("prev2", prev2)
	log.Println("prev1", prev1)

	for range 18 {
		newFib := prev1 + prev2
		log.Println("newFib", newFib)
		prev2 = prev1
		prev1 = newFib
	}
}
