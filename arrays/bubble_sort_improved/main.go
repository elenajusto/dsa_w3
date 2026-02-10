package main

import "log"

/*
If the array is mostly sorted, the previous implementation
will continue to run without swapping elements.

An imrpovement for this is to go through the array and stop
the algorithm if nothing is swapped on first pass.
*/

func main() {
	my_array := []int{7, 3, 9, 12, 11}

	log.Println("initial array:", my_array)
	for index := range len(my_array) { // Start a loop for each item in the array
		swapped := false                                   // Guard condition to check if arrau is already sorted
		for sub_index := range len(my_array) - index - 1 { // Start a loop to compare the above item with every other item in the array
			// Establish the values
			check_1 := &my_array[sub_index]
			check_2 := &my_array[sub_index+1]
			// Check the values
			if *check_1 > *check_2 {
				// Swap if conditon met
				log.Println("bubbling ", *check_1)
				*check_1, *check_2 = *check_2, *check_1
				swapped = false
			}
			log.Println("array after iteration:", my_array)
			if !swapped {
				log.Println("array is already sorted")
				break
			}
		}
	}
	log.Println("final array:", my_array)
}
